package config

import (
	"fmt"
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
	Database        DatabaseConfig
}

// DatabaseConfig holds database connection configuration.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Load reads configuration from environment variables with sensible defaults.
func Load() *Config {
	cfg := &Config{
		Env:             getEnv("ENV", "development"),
		Port:            getEnv("PORT", "8080"),
		DataPath:        getEnv("DATA_PATH", "data/properties.json"),
		CORSAllowOrigin: getEnv("CORS_ALLOW_ORIGIN", "*"),
		LogLevel:        parseLogLevel(getEnv("LOG_LEVEL", "info")),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "property_listing"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	return cfg
}

// DSN returns the database connection string.
func (db DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.DBName, db.SSLMode)
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
