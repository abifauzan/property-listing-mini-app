# Property Listing Mini App — Backend

A RESTful Go API serving property listing data for the DANA take-home assessment.

## Architecture

Clean Architecture with Dependency Injection:

```
cmd/api/          → Application entry point
internal/
  domain/         → Business entities & interfaces
  repository/     → Data access (JSON in-memory store)
  usecase/        → Business logic
  delivery/http/  → HTTP handlers & routing
  middleware/     → CORS & request logging
data/             → JSON data file (mock database)
```

## Prerequisites

- Go 1.22+

## Run the Server

```bash
cd backend
go run cmd/api/main.go
```

The server starts on **http://localhost:8080** by default.

### Environment Variables

| Variable    | Default                  | Description            |
|-------------|--------------------------|------------------------|
| `PORT`      | `8080`                   | Server listen port     |
| `DATA_PATH` | `data/properties.json`   | Path to JSON data file |

## API Endpoints

### List Properties

```
GET /api/v1/properties
GET /api/v1/properties?search={title}
```

Returns all properties, optionally filtered by title (case-insensitive).

### Get Property Detail

```
GET /api/v1/properties/{id}
```

Returns full detail for a single property by its `documentId`.

## Run Tests

```bash
cd backend
go test ./... -v
```
