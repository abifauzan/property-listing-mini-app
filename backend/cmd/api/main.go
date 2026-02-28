package main

import (
	"log/slog"
	"net/http"
	"os"

	delivery "github.com/abifauzan/property-listing-mini-app/backend/internal/delivery/http"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/middleware"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/repository"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/usecase"
)

func main() {
	// Structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// Data file path — configurable via env, defaults to local data directory
	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "data/properties.json"
	}

	// Repository
	repo, err := repository.NewJSONPropertyRepository(dataPath)
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
	app := middleware.Logger(middleware.CORS(mux))

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("server starting", "port", port)
	if err := http.ListenAndServe(":"+port, app); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
