# Property Listing Application — Full Stack Monorepo

A full-stack property listing application with a **Go backend API** and **React Native mobile frontend**.

---

## 📦 Project Structure

```
property-listing-mini-app/
├── backend/                    # Go REST API
│   ├── cmd/api/               # Application entry point
│   ├── internal/              # Clean architecture layers
│   │   ├── domain/           # Business entities
│   │   ├── repository/       # Data access layer
│   │   ├── usecase/          # Business logic
│   │   ├── delivery/         # HTTP handlers
│   │   └── middleware/       # CORS, logging
│   ├── data/                 # JSON data source
│   └── docs/                 # API spec (OpenAPI)
├── mobile/                    # React Native (Expo) app
│   ├── app/                  # Expo Router screens
│   ├── components/           # Reusable UI components
│   ├── hooks/                # React Query hooks
│   ├── services/             # API client
│   ├── types/                # TypeScript types
│   └── utils/                # Utilities
└── miniprogram/              # DANA Mini Program (legacy)
```

---

## 🚀 Quick Start

### Prerequisites

- **Node.js** 18.x or higher
- **Go** 1.21 or higher
- **Expo Go** app (for mobile testing)

### 1. Start the Backend API

```bash
cd backend
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080/api/v1`

### 2. Start the Mobile App

```bash
cd mobile
npm install
npm start
```

Then:
- Press `i` for iOS Simulator
- Press `a` for Android Emulator
- Scan QR code with Expo Go for physical device

---

## 🏗️ Architecture Overview

### Backend (Go)

**Clean Architecture** with dependency injection:

```
HTTP Request
    ↓
Delivery Layer (handlers)
    ↓
Use Case Layer (business logic)
    ↓
Repository Layer (data access)
    ↓
Domain Layer (entities)
```

**Key Features**:
- RESTful API design
- CORS enabled
- JSON file-based data source
- Structured error handling
- Comprehensive logging

### Frontend (React Native)

**Modern React Native** with Expo Router:

```
UI Components
    ↓
React Query Hooks
    ↓
API Services (Axios)
    ↓
Backend API
```

**Key Features**:
- File-based routing (Expo Router)
- TailwindCSS styling (NativeWind)
- Automatic caching (TanStack Query)
- Type-safe API layer
- Animated UI components

---

## 📡 API Endpoints

| Endpoint | Method | Query Params | Description |
|----------|--------|--------------|-------------|
| `/api/v1/properties` | GET | `?search={query}` | Get all properties (with optional search) |
| `/api/v1/properties/{id}` | GET | — | Get property detail by ID |

### Example Requests

```bash
# Get all properties
curl http://localhost:8080/api/v1/properties

# Search properties
curl http://localhost:8080/api/v1/properties?search=villa

# Get property detail
curl http://localhost:8080/api/v1/properties/x0zzhpyrox8ixdy75e2zpw1m
```

---

## 🎯 Features

### Mobile App Features

- ✅ **About Screen** — App information and features
- ✅ **Property Listing** — Browse all properties
- ✅ **Search** — Real-time search with debounce
- ✅ **Layout Toggle** — Switch between list and grid views
- ✅ **Property Detail** — Full property information
- ✅ **Image Gallery** — Carousel with pagination
- ✅ **Pull-to-Refresh** — Reload property data
- ✅ **Loading States** — Animated skeletons
- ✅ **Empty States** — User-friendly messages
- ✅ **Book Now** — Demo booking action

### Backend Features

- ✅ **RESTful API** — Standard HTTP methods
- ✅ **Search Functionality** — Case-insensitive title search
- ✅ **CORS Support** — Cross-origin requests enabled
- ✅ **Error Handling** — Structured error responses
- ✅ **Logging** — Request/response logging
- ✅ **Clean Architecture** — Maintainable codebase

---

## 🛠️ Technology Stack

### Backend

| Component | Technology |
|-----------|-----------|
| Language | Go 1.21+ |
| Framework | Standard library (net/http) |
| Architecture | Clean Architecture |
| Data Source | JSON file |
| API Spec | OpenAPI 3.0 |

### Mobile

| Component | Technology |
|-----------|-----------|
| Framework | Expo SDK 55 |
| Language | TypeScript (strict) |
| Navigation | Expo Router v4 |
| Styling | NativeWind v4 |
| Data Fetching | TanStack Query v5 |
| HTTP Client | Axios |
| Icons | @expo/vector-icons |

---

## 📱 Screenshots

*(Add screenshots after running the app)*

### Mobile App Screens
1. About Screen
2. Property Listing (List View)
3. Property Listing (Grid View)
4. Property Detail
5. Image Gallery Carousel

---

## 🧪 Testing

### Backend Testing

```bash
cd backend
go test ./...
```

### Mobile Testing

Manual testing checklist in `mobile/README.md`

---

## 📚 Documentation

- **Backend API Spec**: `backend/docs/api-spec.yaml` (OpenAPI 3.0)
- **Backend README**: `backend/README.md`
- **Mobile README**: `mobile/README.md`
- **Mobile Summary**: `mobile/PROJECT_SUMMARY.md`

---

## 🔧 Configuration

### Backend Port

Default: `8080`

To change, update `backend/cmd/api/main.go`:

```go
port := ":8080"  // Change this
```

### Mobile API URL

For physical device testing, update `mobile/utils/constants.ts`:

```typescript
export const API_BASE_URL = 'http://YOUR_COMPUTER_IP:8080';
```

---

## 🚢 Deployment

### Backend

```bash
cd backend
go build -o property-api cmd/api/main.go
./property-api
```

### Mobile

```bash
cd mobile
npx eas build --platform ios
npx eas build --platform android
```

*(Requires Expo account and EAS CLI setup)*

---

## 📝 Development Workflow

1. **Start Backend**:
   ```bash
   cd backend && go run cmd/api/main.go
   ```

2. **Start Mobile**:
   ```bash
   cd mobile && npm start
   ```

3. **Make Changes**:
   - Backend: Edit files in `internal/`
   - Mobile: Edit files in `app/`, `components/`, etc.

4. **Test**:
   - Backend: `go test ./...`
   - Mobile: Manual testing via Expo Go

---

## 🎨 Design Decisions

### Why Go for Backend?

- Fast compilation and execution
- Strong standard library
- Excellent for REST APIs
- Clean architecture support
- Type safety

### Why Expo for Mobile?

- Fast development iteration
- OTA updates
- Cross-platform (iOS + Android)
- Modern routing (Expo Router)
- Rich ecosystem

### Why Clean Architecture?

- Separation of concerns
- Testability
- Maintainability
- Scalability
- Technology independence

---

## 🔄 Future Enhancements

### Backend
- [ ] Database integration (PostgreSQL)
- [ ] Authentication & authorization
- [ ] Pagination support
- [ ] Filtering by price, facilities
- [ ] Image upload service
- [ ] Unit tests & integration tests

### Mobile
- [ ] Favorites functionality
- [ ] Advanced filtering
- [ ] Map view
- [ ] User authentication
- [ ] Booking history
- [ ] Push notifications
- [ ] Offline mode

---

## 🐛 Troubleshooting

### "Connection refused" on mobile

- Ensure backend is running
- For physical devices, use computer's IP instead of `localhost`
- Check firewall settings

### TypeScript errors in mobile

```bash
cd mobile
npm install
```

### Backend port already in use

```bash
lsof -ti:8080 | xargs kill -9
```

---

## 📄 License

MIT

---

## 👤 Author

**Abi Fauzan**  
Email: abifauzan234@gmail.com

---

## 🙏 Acknowledgments

Built as part of the DANA technical assessment.

**Tech Stack**:
- Go standard library
- Expo & React Native
- TanStack Query
- NativeWind
- Axios
