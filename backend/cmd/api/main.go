package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Get property count for health check
	properties, _ := repo.GetAll(context.Background())
	propertyCount := len(properties)

	// Usecase
	uc := usecase.NewPropertyUsecase(repo)

	// HTTP handler & routes
	mux := http.NewServeMux()

	// Property endpoints
	propertyHandler := delivery.NewPropertyHandler(uc)
	propertyHandler.RegisterRoutes(mux)

	// Health check endpoints
	healthHandler := delivery.NewHealthHandler(true, propertyCount)
	healthHandler.RegisterRoutes(mux)

	// Middleware chain: RequestID -> Timeout -> Logger -> CORS -> Router
	app := middleware.RequestID(
		middleware.Timeout(30 * time.Second)(
			middleware.Logger(
				middleware.CORS(mux),
			),
		),
	)

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      app,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("server starting", "port", port, "properties", propertyCount)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-shutdownChan
	slog.Info("shutdown signal received, gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped gracefully")
}
