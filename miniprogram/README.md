# Property Finder — DANA Mini Program Frontend

A property listing mini-app built with the **DANA Mini Program Framework**. Browse properties, search by title, toggle between list/grid views, and view detailed property information.

## Prerequisites

- **DANA Mini Program IDE** (download from [mini-program.dana.id](https://mini-program.dana.id))
- **Go backend** running on `http://localhost:8080` (see `../backend/README.md`)

## Getting Started

### 1. Start the Backend

```bash
cd ../backend
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080/api/v1`.

### 2. Open in DANA Mini Program IDE

1. Open the DANA Mini Program IDE.
2. Select **Open Project** and navigate to the `miniprogram/` directory.
3. The IDE will detect `app.json` and load the project.
4. Click **Preview** or **Simulate** to run the app.

### 3. Regenerate Icons (Optional)

If you need to regenerate the icon assets:

```bash
python3 scripts/generate-icons.py
```

## Project Structure

```
miniprogram/
├── app.js                          # App entry point & global config
├── app.json                        # App routes, tab bar, window settings
├── app.acss                        # Global styles & utility classes
├── pages/
│   ├── about/                      # Static "About" page
│   ├── listing/                    # Property listing with search & toggle
│   └── detail/                     # Property detail with gallery & booking
├── components/
│   ├── property-card/              # List-view property card
│   ├── property-grid-item/         # Grid-view property card
│   ├── search-bar/                 # Search input with clear button
│   ├── layout-toggle/              # List/grid view toggle button
│   ├── loading-skeleton/           # Shimmer loading placeholders
│   ├── empty-state/                # Empty/error state display
│   └── image-carousel/             # Swiper-based image gallery
├── services/
│   ├── property.service.js         # Business logic, caching, data transform
│   └── storage.service.js          # Local storage management
├── api/
│   ├── config.js                   # API configuration
│   ├── http.client.js              # HTTP client with retry & error handling
│   └── property.api.js             # Property API endpoints
├── utils/
│   ├── constants.js                # App-wide constants
│   ├── format.js                   # Price & date formatting
│   └── debounce.js                 # Debounce utility
└── assets/images/                  # Icons & logo
```

## Architecture

The app follows **Clean Architecture** with clear layer separation:

```
Pages (Presentation) → Components (UI) → Services (Business Logic) → API (Data Access)
```

- **Pages**: Manage page state, lifecycle, and user interactions
- **Components**: Reusable, props-driven, stateless UI elements
- **Services**: Data fetching, caching, transformation, and filtering
- **API Layer**: HTTP client with retry logic, error parsing, and response mapping

## Features

| Feature | Description |
|---|---|
| **Tab Navigation** | Bottom tab bar with About and Listings tabs |
| **Property Listing** | Fetches and displays properties from backend API |
| **Search** | Debounced (300ms) client-side search filtering by title |
| **Layout Toggle** | Switch between vertical list and 2-column grid views |
| **View Persistence** | Layout preference saved to local storage |
| **Property Detail** | Full details with hero image, gallery, facilities, terms |
| **Image Carousel** | Swiper with counter indicator and full-screen preview |
| **Book Now** | Dummy booking action with toast notification |
| **Loading Skeletons** | Shimmer placeholders for list, grid, and detail views |
| **Empty States** | Helpful messages for no results and errors |
| **Pull to Refresh** | Refresh listing data with cache invalidation |
| **Error Handling** | Network errors, 404s, and server errors with retry |
| **Caching** | 5-minute TTL cache for listings and detail data |

## API Integration

The frontend connects to the Go backend at `http://localhost:8080/api/v1`:

| Endpoint | Method | Description |
|---|---|---|
| `/properties` | GET | List properties (optional `?search=` query) |
| `/properties/{id}` | GET | Get property detail by document ID |

To change the API base URL, edit `app.js`:

```javascript
App({
  globalData: {
    apiBaseURL: 'http://your-api-url/api/v1',
  },
});
```

## Configuration

| Setting | File | Default |
|---|---|---|
| API Base URL | `app.js` | `http://localhost:8080/api/v1` |
| Request Timeout | `api/config.js` | 10,000ms |
| Retry Attempts | `api/config.js` | 2 |
| Cache TTL | `utils/constants.js` | 5 minutes |
| Search Debounce | `utils/constants.js` | 300ms |
