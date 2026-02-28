# Property Listing Mini App API Documentation

## Overview

The Property Listing Mini App API provides RESTful endpoints for serving property data to a mini-program frontend. Built with Go using clean architecture principles, this API offers property listing retrieval with search functionality and detailed property information.

**Base URL**: `http://localhost:8080` (development)  
**API Version**: v1  
**Content-Type**: `application/json`

---

## Endpoints

### 1. Get Property Listings

Retrieve a list of properties with optional search filtering.

**Endpoint**: `GET /api/v1/properties`

#### Parameters

| Name | Type | Required | Description | Example |
|------|------|----------|-------------|---------|
| search | string | No | Search term to filter properties by title (case-insensitive) | `villa` |

#### Response Format

```json
{
  "data": {
    "propertyListings": [
      {
        "documentId": "x0zzhpyrox8ixdy75e2zpw1m",
        "Banner": {
          "url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg"
        },
        "Title": "My Villa Sample",
        "Price": "750000",
        "createdAt": "2025-01-31T08:18:56.778Z"
      }
    ]
  }
}
```

#### Examples

**Get all properties:**
```bash
curl -X GET "http://localhost:8080/api/v1/properties"
```

**Search properties by title:**
```bash
curl -X GET "http://localhost:8080/api/v1/properties?search=villa"
```

---

### 2. Get Property Detail

Retrieve detailed information for a specific property by its document ID.

**Endpoint**: `GET /api/v1/properties/{id}`

#### Parameters

| Name | Type | Required | Description | Example |
|------|------|----------|-------------|---------|
| id | string | Yes | Property document ID | `x0zzhpyrox8ixdy75e2zpw1m` |

#### Response Format

```json
{
  "data": {
    "propertyListings": [
      {
        "documentId": "x0zzhpyrox8ixdy75e2zpw1m",
        "Banner": {
          "url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg"
        },
        "Title": "My Villa Sample",
        "Price": "750000",
        "createdAt": "2025-01-31T08:18:56.778Z",
        "Description": "Aenean tellus velit, porttitor eget diam vel, pharetra pellentesque ante...",
        "Images": [
          {
            "url": "https://res.cloudinary.com/dzktit9nc/image/upload/v1733987640/pexels_eray_ozdogan_615189320_18785790_736c45bfb0.jpg"
          }
        ],
        "Facilities": ["Kitchen", "Free Park", "Bar"],
        "Terms": "Donec gravida, quam at volutpat pharetra, elit ante ullamcorper eros...",
        "Conditions": "Vestibulum sit amet vehicula tellus. Etiam sed pharetra risus..."
      }
    ]
  }
}
```

#### Example

```bash
curl -X GET "http://localhost:8080/api/v1/properties/x0zzhpyrox8ixdy75e2zpw1m"
```

---

## Data Models

### PropertyListing (Minimal)

Used in the listings endpoint for optimal performance.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| documentId | string | Yes | Unique property identifier |
| Banner | object | Yes | Banner image with URL |
| Title | string | Yes | Property title |
| Price | string | Yes | Property price (numeric string) |
| createdAt | string | Yes | ISO 8601 timestamp |

### Property (Complete)

Used in the detail endpoint with all property information.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| documentId | string | Yes | Unique property identifier |
| Banner | object | Yes | Banner image with URL |
| Title | string | Yes | Property title |
| Price | string | Yes | Property price (numeric string) |
| createdAt | string | Yes | ISO 8601 timestamp |
| Description | string | Yes | Detailed property description |
| Images | array | Yes | Gallery of property images |
| Facilities | array | Yes | Available facilities |
| Terms | string | Yes | Property terms and conditions |
| Conditions | string | Yes | Additional conditions |

### Banner

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| url | string | Yes | Image URL |

### Image

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| url | string | Yes | Image URL |

---

## Error Responses

All error responses follow this format:

```json
{
  "error": {
    "message": "Error description",
    "code": "ERROR_CODE"
  }
}
```

### Common HTTP Status Codes

| Status | Description | Example |
|--------|-------------|---------|
| 200 | Success | - |
| 400 | Bad Request | Invalid parameters |
| 404 | Not Found | Property not found |
| 500 | Internal Server Error | Server error |

---

## Features

### Search Functionality
- **Case-insensitive**: Search ignores letter case
- **Title-based**: Searches only in property titles
- **Partial matching**: Finds properties containing the search term
- **Optional**: Can be omitted to get all properties

### Response Structure
- **Consistent envelope**: All responses wrapped in `data` object
- **Array format**: Properties always returned in `propertyListings` array
- **Schema compliance**: Follows provided JSON schema requirements

---

## Usage Examples

### JavaScript/TypeScript

```typescript
// Get all properties
const response = await fetch('http://localhost:8080/api/v1/properties');
const data = await response.json();

// Search properties
const searchResponse = await fetch('http://localhost:8080/api/v1/properties?search=villa');
const searchData = await searchResponse.json();

// Get property detail
const detailResponse = await fetch('http://localhost:8080/api/v1/properties/x0zzhpyrox8ixdy75e2zpw1m');
const detailData = await detailResponse.json();
```

### Python

```python
import requests

# Get all properties
response = requests.get('http://localhost:8080/api/v1/properties')
data = response.json()

# Search properties
search_response = requests.get('http://localhost:8080/api/v1/properties', 
                              params={'search': 'villa'})
search_data = search_response.json()

# Get property detail
detail_response = requests.get('http://localhost:8080/api/v1/properties/x0zzhpyrox8ixdy75e2zpw1m')
detail_data = detail_response.json()
```

---

## Architecture Notes

- **Clean Architecture**: Separated into domain, usecase, repository, and delivery layers
- **JSON File Storage**: Uses JSON file for data persistence (easily migratable to database)
- **CORS Enabled**: Supports cross-origin requests for mini-program frontend
- **Structured Logging**: Uses JSON structured logging for better observability
- **Error Handling**: Comprehensive error handling with proper HTTP status codes

---

## Development

### Running the Server

```bash
cd backend
go run cmd/api/main.go
```

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| PORT | 8080 | Server port |
| DATA_PATH | data/properties.json | Path to JSON data file |

### Testing

```bash
go test ./...
```

---

## OpenAPI Specification

A complete OpenAPI 3.0 specification is available at `docs/api-spec.yaml` for import into API documentation tools like Swagger UI, Postman, or Insomnia.
