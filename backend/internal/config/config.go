package config

import (
	"log/slog"
	"os"
	"strings"
)

// Config holds all application configuration loaded from environment variables.
type Config struct {
	Env             string
	Port            string
	DataPath        string
	CORSAllowOrigin string
	LogLevel        slog.Level
}

// Load reads configuration from environment variables with sensible defaults.
func Load() *Config {
	cfg := &Config{
		Env:             getEnv("ENV", "development"),
		Port:            getEnv("PORT", "8080"),
		DataPath:        getEnv("DATA_PATH", "data/properties.json"),
		CORSAllowOrigin: getEnv("CORS_ALLOW_ORIGIN", "*"),
		LogLevel:        parseLogLevel(getEnv("LOG_LEVEL", "info")),
	}

	return cfg
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// parseLogLevel converts a string log level to slog.Level.
func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
