package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/abifauzan/property-listing-mini-app/backend/internal/config"
	delivery "github.com/abifauzan/property-listing-mini-app/backend/internal/delivery/http"
	"github.com/abifauzan/property-listing-mini-app/backend/internal/domain"
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

	// Repository - choose between PostgreSQL and JSON based on configuration
	var repo domain.PropertyRepository
	var err error

	if usePostgres(cfg) {
		slog.Info("initializing PostgreSQL repository")
		repo, err = repository.NewPostgresPropertyRepository(cfg.Database.DSN())
		if err != nil {
			slog.Error("failed to initialize PostgreSQL repository", "error", err)
			os.Exit(1)
		}
	} else {
		slog.Info("initializing JSON repository", "data_path", cfg.DataPath)
		repo, err = repository.NewJSONPropertyRepository(cfg.DataPath)
		if err != nil {
			slog.Error("failed to initialize JSON repository", "error", err)
			os.Exit(1)
		}
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

// usePostgres determines whether to use PostgreSQL repository based on configuration.
// Returns true if DATABASE_URL is set or all database config variables are present.
func usePostgres(cfg *config.Config) bool {
	// Check if DATABASE_URL environment variable is set
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return true
	}

	// Check if all required database config is available (non-default)
	return cfg.Database.Host != "localhost" ||
		cfg.Database.User != "postgres" ||
		cfg.Database.Password != "password" ||
		cfg.Database.DBName != "property_listing" ||
		cfg.Database.Port != "5432"
}
