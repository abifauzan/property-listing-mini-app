# Property Listing Backend — Go REST API

A production-ready RESTful API built with **Go** using **Clean Architecture** principles, serving property listing data for the DANA Mini Program frontend.

---

## 🏗️ Architecture

### Clean Architecture Layers

```
cmd/api/                    → Application entry point & dependency wiring
internal/
  ├── domain/              → Business entities & repository interfaces
  ├── repository/          → Data access implementation (JSON file)
  ├── usecase/             → Business logic & orchestration
  ├── delivery/http/       → HTTP handlers, routing, request/response
  └── middleware/          → Cross-cutting concerns (CORS, logging)
data/                      → JSON data source (properties.json)
docs/                      → API documentation & OpenAPI spec
```

### Dependency Flow

```
HTTP Request
    ↓
Middleware (CORS, Logging)
    ↓
Delivery Layer (handlers)
    ↓
Use Case Layer (business logic)
    ↓
Repository Layer (data access)
    ↓
Domain Layer (entities)
```

**Key Principles**:
- **Dependency Inversion**: Inner layers define interfaces, outer layers implement them
- **Single Responsibility**: Each layer has one clear purpose
- **Testability**: All layers are independently testable via interfaces
- **Separation of Concerns**: Business logic isolated from infrastructure

---

## 🚀 Quick Start

### Prerequisites

- **Go** 1.25.0 or higher
- **Git** (for cloning)

### Installation & Run

```bash
# Navigate to backend directory
cd backend

# Run the server
go run cmd/api/main.go
```

The server will start on **`http://localhost:8080`**

### Verify Server is Running

```bash
# Test health check
curl http://localhost:8080/api/v1/properties

# Expected response: JSON with property listings
```

---

## 📡 API Endpoints

### 1. List Properties

**Endpoint**: `GET /api/v1/properties`

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `search` | string | No | Filter properties by title (case-insensitive) |

**Response**:
```json
{
  "data": {
    "propertyListings": [
      {
        "documentId": "x0zzhpyrox8ixdy75e2zpw1m",
        "Banner": { "url": "https://..." },
        "Title": "My Villa Sample",
        "Price": "750000",
        "createdAt": "2025-01-31T08:18:56.778Z"
      }
    ]
  }
}
```

**Examples**:
```bash
# Get all properties
curl http://localhost:8080/api/v1/properties

# Search by title
curl http://localhost:8080/api/v1/properties?search=villa
```

---

### 2. Get Property Detail

**Endpoint**: `GET /api/v1/properties/{id}`

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `id` | string | Yes | Property document ID |

**Response**:
```json
{
  "data": {
    "propertyListings": [
      {
        "documentId": "x0zzhpyrox8ixdy75e2zpw1m",
        "Banner": { "url": "https://..." },
        "Title": "My Villa Sample",
        "Price": "750000",
        "Description": "Full description...",
        "Images": [{ "url": "https://..." }],
        "Facilities": ["Kitchen", "Free Park"],
        "Terms": "Terms and conditions...",
        "Conditions": "Additional conditions...",
        "createdAt": "2025-01-31T08:18:56.778Z"
      }
    ]
  }
}
```

**Example**:
```bash
curl http://localhost:8080/api/v1/properties/x0zzhpyrox8ixdy75e2zpw1m
```

---

## ⚙️ Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server listen port |
| `DATA_PATH` | `data/properties.json` | Path to JSON data file |

**Example**:
```bash
# Custom port
PORT=3000 go run cmd/api/main.go

# Custom data path
DATA_PATH=/path/to/data.json go run cmd/api/main.go
```

---

## 🧪 Testing

### Run All Tests

```bash
cd backend
go test ./... -v
```

### Run Tests with Coverage

```bash
go test ./... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Structure

- **Repository Tests** (`internal/repository/json_repository_test.go`)
  - Test JSON file loading
  - Test search functionality
  - Test property retrieval by ID
  
- **Use Case Tests** (`internal/usecase/property_usecase_test.go`)
  - Test business logic
  - Mock repository for isolation
  - Test error handling

---

## 📂 Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   ├── domain/
│   │   └── property.go               # Entity & repository interface
│   ├── repository/
│   │   ├── json_repository.go        # JSON file implementation
│   │   └── json_repository_test.go   # Repository tests
│   ├── usecase/
│   │   ├── property_usecase.go       # Business logic
│   │   └── property_usecase_test.go  # Use case tests
│   ├── delivery/
│   │   └── http/
│   │       ├── handler.go            # HTTP handlers
│   │       └── response.go           # Response helpers
│   └── middleware/
│       ├── cors.go                   # CORS middleware
│       └── logger.go                 # Request logging
├── data/
│   └── properties.json               # Property data source
├── docs/
│   ├── README.md                     # API documentation
│   └── api-spec.yaml                 # OpenAPI 3.0 spec
├── go.mod                            # Go module definition
└── README.md                         # This file
```

---

## 🛠️ Technology Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Language | Go 1.25.0 | Backend implementation |
| HTTP Server | `net/http` (stdlib) | HTTP server & routing |
| Logging | `log/slog` (stdlib) | Structured JSON logging |
| Data Storage | JSON file | Mock database |
| Testing | `testing` (stdlib) | Unit & integration tests |
| Architecture | Clean Architecture | Maintainable, testable code |

**No External Dependencies**: This project uses only Go's standard library for maximum simplicity and reliability.

---

## 🔍 Code Quality

### Design Patterns

- **Repository Pattern**: Abstracts data access
- **Dependency Injection**: Constructor-based DI
- **Interface Segregation**: Small, focused interfaces
- **Error Wrapping**: Contextual error messages

### Best Practices

- ✅ Structured logging with `slog`
- ✅ Graceful error handling
- ✅ CORS enabled for cross-origin requests
- ✅ Request/response logging
- ✅ Interface-based design for testability
- ✅ No global state
- ✅ Idiomatic Go code

---

## 🚢 Deployment

### Build Binary

```bash
cd backend
go build -o property-api cmd/api/main.go
```

### Run Binary

```bash
./property-api
```

### Docker Deployment (Example)

```dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o property-api cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/property-api .
COPY --from=builder /app/data ./data
EXPOSE 8080
CMD ["./property-api"]
```

---

## 📚 Documentation

- **API Documentation**: [`docs/README.md`](docs/README.md)
- **OpenAPI Spec**: [`docs/api-spec.yaml`](docs/api-spec.yaml)
- **Root README**: [`../README.md`](../README.md)

---

## 🐛 Troubleshooting

### Port Already in Use

```bash
# Find process using port 8080
lsof -ti:8080

# Kill the process
lsof -ti:8080 | xargs kill -9
```

### Data File Not Found

Ensure `data/properties.json` exists:
```bash
ls -la data/properties.json
```

Or set custom path:
```bash
DATA_PATH=/path/to/data.json go run cmd/api/main.go
```

### CORS Issues

CORS is enabled by default for all origins. To restrict:
Edit `internal/middleware/cors.go` and modify allowed origins.

---

## 🔄 Future Enhancements

- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] Authentication & authorization (JWT)
- [ ] Pagination support
- [ ] Advanced filtering (price range, facilities)
- [ ] Rate limiting
- [ ] API versioning
- [ ] Metrics & monitoring (Prometheus)
- [ ] Graceful shutdown
- [ ] Health check endpoint
- [ ] Docker Compose setup

---

## 📄 License

MIT

---

## 👤 Author

**Abi Fauzan**  
Email: abifauzan234@gmail.com

---

## 🙏 Acknowledgments

Built as part of the **DANA Technical Assessment** demonstrating:
- Clean Architecture in Go
- RESTful API design
- Test-driven development
- Production-ready code patterns
