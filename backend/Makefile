.PHONY: run run-dev run-staging run-prod migrate-up migrate-down seed-db help

# Default target: run in development mode
run: run-dev

# Run in development mode (localhost)
run-dev:
	@echo "Starting server in development mode..."
	ENV=development \
	PORT=8080 \
	DATA_PATH=data/properties.json \
	CORS_ALLOW_ORIGIN=* \
	LOG_LEVEL=info \
	DB_HOST=localhost DB_PORT=5432 DB_USER=abifauzan DB_PASSWORD="" DB_NAME=property_listing DB_SSLMODE=disable \
	go run cmd/api/main.go

# Run in staging mode
run-staging:
	@echo "Starting server in staging mode..."
	ENV=staging \
	PORT=8080 \
	DATA_PATH=data/properties.json \
	CORS_ALLOW_ORIGIN=https://staging.example.com \
	LOG_LEVEL=info \
	DB_HOST=localhost DB_PORT=5432 DB_USER=abifauzan DB_PASSWORD="" DB_NAME=property_listing DB_SSLMODE=disable \
	go run cmd/api/main.go

# Run in production mode
run-prod:
	@echo "Starting server in production mode..."
	ENV=production \
	PORT=8080 \
	DATA_PATH=data/properties.json \
	CORS_ALLOW_ORIGIN=https://example.com \
	LOG_LEVEL=warn \
	DB_HOST=localhost DB_PORT=5432 DB_USER=abifauzan DB_PASSWORD="" DB_NAME=property_listing DB_SSLMODE=disable \
	go run cmd/api/main.go

# Database migrations
migrate-up:
	@echo "Running database migrations..."
	DB_HOST=localhost DB_PORT=5432 DB_USER=abifauzan DB_PASSWORD="" DB_NAME=property_listing DB_SSLMODE=disable go run cmd/migrate/main.go db-up

migrate-down:
	@echo "Rolling back database migrations..."
	DB_HOST=localhost DB_PORT=5432 DB_USER=abifauzan DB_PASSWORD="" DB_NAME=property_listing DB_SSLMODE=disable go run cmd/migrate/main.go db-down

seed-db:
	@echo "Seeding database with JSON data..."
	DB_HOST=localhost DB_PORT=5432 DB_USER=abifauzan DB_PASSWORD="" DB_NAME=property_listing DB_SSLMODE=disable go run cmd/migrate/main.go seed

# Show available commands
help:
	@echo "Available targets:"
	@echo "  make run          - Run in development mode (default)"
	@echo "  make run-dev      - Run in development mode"
	@echo "  make run-staging  - Run in staging mode"
	@echo "  make run-prod     - Run in production mode"
	@echo "  make migrate-up   - Run database migrations"
	@echo "  make migrate-down - Rollback database migrations"
	@echo "  make seed-db      - Seed database with JSON data"
	@echo "  make help         - Show this help message"
