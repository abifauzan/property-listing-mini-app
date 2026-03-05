package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/config"
	delivery "github.com/abifauzan/property-listing-mini-app/backend/internal/delivery/http"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/middleware"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/repository"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/usecase"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Structured logger with configured level
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))
	slog.SetDefault(logger)

	slog.Info("configuration loaded", "env", cfg.Env, "port", cfg.Port)

	// Repository
	repo, err := repository.NewJSONPropertyRepository(cfg.DataPath)
	if err != nil {
		slog.Error("failed to initialize repository", "error", err)
		os.Exit(1)
	}

	// Usecase
	uc := usecase.NewPropertyUsecase(repo)

	// HTTP handler & routes
	mux := http.NewServeMux()
	handler := delivery.NewPropertyHandler(uc)
	handler.RegisterRoutes(mux)

	// Middleware chain: Logger -> CORS -> Router
	app := middleware.Logger(middleware.CORS(mux, cfg.CORSAllowOrigin))

	// Server
	slog.Info("server starting", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, app); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
