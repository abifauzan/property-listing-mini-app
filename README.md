# DANA Mini Program Property Listing

A property listing mini application built using the DANA Mini Program framework.

## Project Structure

```
property-listing-mini-app/
├── backend/                 # Go backend API
├── frontend/               # DANA Mini Program frontend
├── docs/                   # Documentation
└── README.md              # This file
```

## Git Branching Strategy

This project follows a structured git branching workflow. See [Git Branching Rules](docs/git-branching-rules.md) for detailed guidelines.

### Current Branches
- **Release**: `release/2026_Mar_dana_mini_app`
- **Feature**: `feature/DANA-003-frontend-implementation`

## Getting Started

### Backend Setup
```bash
cd backend
go mod download
go run cmd/api/main.go
```

### Frontend Setup
```bash
cd frontend
# Follow DANA Mini Program IDE setup instructions
```

## Documentation

- [API Documentation](backend/docs/api-spec.yaml)
- [Git Branching Rules](docs/git-branching-rules.md)

## Requirements

- DANA Mini Program Framework
- Go 1.21+
- DANA Mini Program IDE
