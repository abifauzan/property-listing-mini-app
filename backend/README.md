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
- **PostgreSQL** (optional, for database mode)

### Installation & Run

#### Option 1: JSON File Mode (Default)

```bash
# Navigate to backend directory
cd backend

# Run the server with JSON data source
go run cmd/api/main.go
```

#### Option 2: PostgreSQL Database Mode

```bash
# 1. Install PostgreSQL
# On macOS with Homebrew:
brew install postgresql
brew services start postgresql

# 2. Create database
createdb property_listing

# 3. Run migrations
make migrate-up

# 4. Seed database with existing JSON data
make seed-db

# 5. Run server with database
DB_HOST=localhost DB_USER=postgres DB_PASSWORD=password go run cmd/api/main.go
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
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | PostgreSQL username |
| `DB_PASSWORD` | `password` | PostgreSQL password |
| `DB_NAME` | `property_listing` | PostgreSQL database name |
| `DB_SSLMODE` | `disable` | PostgreSQL SSL mode |
| `DATABASE_URL` | - | Full database URL (overrides individual DB vars) |

**Examples**:
```bash
# Custom port
PORT=3000 go run cmd/api/main.go

# Custom data path
DATA_PATH=/path/to/data.json go run cmd/api/main.go

# Database mode
DATABASE_URL=postgres://user:pass@host:5432/db go run cmd/api/main.go

# Individual database variables
DB_HOST=localhost DB_USER=postgres DB_PASSWORD=mypassword go run cmd/api/main.go
```

**Repository Selection**: The application automatically chooses between JSON and PostgreSQL repositories:
- Uses **PostgreSQL** if `DATABASE_URL` is set or database config differs from defaults
- Uses **JSON** if database config is at default values

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

- **Repository Tests** (`internal/repository/`)
  - `json_repository_test.go` - Test JSON file loading
  - `postgres_repository_test.go` - Test PostgreSQL operations
  - Test search functionality
  - Test property retrieval by ID
  
- **Use Case Tests** (`internal/usecase/property_usecase_test.go`)
  - Test business logic
  - Mock repository for isolation
  - Test error handling

### Database Testing

PostgreSQL tests require a test database:
```bash
# Create test database
createdb property_listing_test

# Run tests with database
SKIP_DB_TESTS= go test ./... -v

# Skip database tests
SKIP_DB_TESTS=1 go test ./... -v
```

---

## 📂 Project Structure

```
backend/
├── cmd/
│   ├── api/
│   │   └── main.go                       # Application entry point
│   └── migrate/
│       └── main.go                       # Database migration tool
├── internal/
│   ├── config/
│   │   └── config.go                     # Configuration management
│   ├── domain/
│   │   └── property.go                  # Entity & repository interface
│   ├── migrate/
│   │   └── migrate.go                    # Migration utilities
│   ├── repository/
│   │   ├── json_repository.go           # JSON file implementation
│   │   ├── postgres_repository.go       # PostgreSQL implementation
│   │   ├── json_repository_test.go      # JSON repository tests
│   │   └── postgres_repository_test.go   # PostgreSQL repository tests
│   ├── usecase/
│   │   ├── property_usecase.go          # Business logic
│   │   └── property_usecase_test.go     # Use case tests
│   ├── delivery/
│   │   └── http/
│   │       ├── property_handler.go      # HTTP handlers
│   │       └── property_handler_test.go # Handler tests
│   └── middleware/
│       └── middleware.go                # CORS & logging middleware
├── migrations/                          # Database migration files
│   ├── 000001_create_properties_table.up.sql
│   ├── 000001_create_properties_table.down.sql
│   ├── 000002_create_images_table.up.sql
│   ├── 000002_create_images_table.down.sql
│   ├── 000003_create_facilities_table.up.sql
│   └── 000003_create_facilities_table.down.sql
├── data/
│   └── properties.json                  # Property data source
├── docs/
│   ├── README.md                        # API documentation
│   └── api-spec.yaml                    # OpenAPI 3.0 spec
├── Makefile                             # Build & migration commands
├── go.mod                               # Go module definition
└── README.md                            # This file
```

---

## 🛠️ Technology Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Language | Go 1.25.0 | Backend implementation |
| HTTP Server | `net/http` (stdlib) | HTTP server & routing |
| Database | PostgreSQL (optional) | Persistent data storage |
| Database Driver | `pgx/v5` | PostgreSQL connection |
| Migrations | `golang-migrate/migrate` | Database schema management |
| Logging | `log/slog` (stdlib) | Structured JSON logging |
| Data Storage | JSON file | Fallback/mock database |
| Testing | `testing` (stdlib) | Unit & integration tests |
| Architecture | Clean Architecture | Maintainable, testable code |

**Dependencies**: Minimal external dependencies for maximum reliability. PostgreSQL is optional - the application works perfectly with JSON files.

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

## 🔄 Database Management

### Make Commands

```bash
# Database migrations
make migrate-up    # Run all migrations
make migrate-down  # Rollback all migrations
make seed-db       # Seed database with JSON data

# Application
make run          # Run in development mode
make run-dev      # Run in development mode
make run-staging  # Run in staging mode
make run-prod     # Run in production mode
make help         # Show all commands
```

### Database Schema

The PostgreSQL implementation uses three tables:

1. **properties** - Main property data
2. **property_images** - Property gallery images (one-to-many)
3. **property_facilities** - Property facility tags (many-to-many)

All tables are properly indexed for performance and include foreign key constraints.

### Migration Files

Migration files are located in the `migrations/` directory:
- `*.up.sql` - Apply migration
- `*.down.sql` - Rollback migration

Migrations are tracked in the `schema_migrations` table.

---

## 🔄 Future Enhancements

- [x] ~~Database integration (PostgreSQL/MongoDB)~~ ✅ **COMPLETED**
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
