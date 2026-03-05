.PHONY: run run-dev run-staging run-prod help

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
	go run cmd/api/main.go

# Run in staging mode
run-staging:
	@echo "Starting server in staging mode..."
	ENV=staging \
	PORT=8080 \
	DATA_PATH=data/properties.json \
	CORS_ALLOW_ORIGIN=https://staging.example.com \
	LOG_LEVEL=info \
	go run cmd/api/main.go

# Run in production mode
run-prod:
	@echo "Starting server in production mode..."
	ENV=production \
	PORT=8080 \
	DATA_PATH=data/properties.json \
	CORS_ALLOW_ORIGIN=https://example.com \
	LOG_LEVEL=warn \
	go run cmd/api/main.go

# Show available commands
help:
	@echo "Available targets:"
	@echo "  make run          - Run in development mode (default)"
	@echo "  make run-dev      - Run in development mode"
	@echo "  make run-staging  - Run in staging mode"
	@echo "  make run-prod     - Run in production mode"
	@echo "  make help         - Show this help message"
