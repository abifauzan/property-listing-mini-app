# Property Listing Frontend — DANA Mini Program

A modern, responsive property listing application built with the **DANA Mini Program Framework**, featuring real-time search, layout toggling, and smooth user experience.

---

## 🏗️ Architecture

### Component-Based Structure

```
frontend/
├── pages/                      # Application screens
│   ├── about/                 # About page (static info)
│   ├── listing/               # Property listing with search
│   ├── detail/                # Property detail view
│   └── index/                 # Landing page
├── components/                # Reusable UI components
│   ├── search-bar/           # Search input component
│   ├── layout-toggle/        # View mode switcher
│   ├── property-card/        # List view card
│   ├── property-grid-item/   # Grid view card
│   ├── image-carousel/       # Image gallery carousel
│   ├── loading-skeleton/     # Loading state UI
│   └── empty-state/          # Empty state UI
├── services/                  # Business logic layer
│   ├── property.service.js   # Property data operations
│   └── storage.service.js    # Local storage management
├── api/                       # HTTP client layer
│   ├── config.js             # API configuration
│   ├── http.client.js        # HTTP request wrapper
│   └── property.api.js       # Property API endpoints
├── utils/                     # Helper functions
│   ├── constants.js          # App constants
│   ├── debounce.js           # Debounce utility
│   └── formatter.js          # Data formatters
├── assets/                    # Static files
│   └── images/               # Icons and images
├── app.js                     # Application entry point
├── app.json                   # App configuration
└── app.acss                   # Global styles
```

### Data Flow

```
User Interaction
    ↓
Page (AXML + JS)
    ↓
Component (reusable UI)
    ↓
Service (business logic)
    ↓
API Client (HTTP layer)
    ↓
Backend API
```

---

## 🚀 Quick Start

### Prerequisites

- **DANA Mini Program IDE** (latest version)
- **Backend API** running on `http://localhost:8080`

### Installation

1. **Open DANA Mini Program IDE**
2. **Import Project**:
   - Click **File** → **Import Project**
   - Select the `frontend/` directory
   - Click **Open**
3. **Run Simulator**:
   - Click the **Run** button
   - The simulator will launch with the app

### Configuration

Update API endpoint in `api/config.js`:

```javascript
const config = {
  get baseURL() {
    return 'http://localhost:8080/api/v1';
  },
  timeout: 10000,
};
```

For physical device testing, use your computer's IP:

```javascript
const config = {
  get baseURL() {
    return 'http://192.168.1.100:8080/api/v1';
  },
};
```

---

## 📱 Features

### ✅ Implemented Features

#### About Page
- Static information page
- App description and features
- Clean, modern UI

#### Property Listing Page
- **Search Bar**: Real-time property search by title
- **Layout Toggle**: Switch between list and grid views
- **Pull-to-Refresh**: Reload property data
- **Loading States**: Skeleton screens during data fetch
- **Empty States**: User-friendly messages when no results
- **Persistent View Mode**: Remembers user's layout preference

#### Property Detail Page
- **Full Property Information**: Title, price, description
- **Image Carousel**: Swipeable gallery with pagination dots
- **Facilities List**: Display property amenities
- **Terms & Conditions**: Show property terms
- **Book Now Button**: Demo booking action (shows toast)

#### Components
- **Search Bar**: Debounced search input
- **Layout Toggle**: List/Grid view switcher
- **Property Card**: List view card with image, title, price
- **Property Grid Item**: Grid view card (compact)
- **Image Carousel**: Touch-enabled image gallery
- **Loading Skeleton**: Animated loading placeholders
- **Empty State**: No results or error messages

---

## 🛠️ Technology Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Framework | DANA Mini Program | Mini-program runtime |
| Markup | AXML (Ant XML) | UI templates |
| Styling | ACSS (Ant CSS) | Component styling |
| Scripting | JavaScript (ES6+) | Business logic |
| HTTP Client | `my.request` | API communication |
| State Management | `setData()` | Page state updates |
| Storage | `my.setStorage` | Local data persistence |
| Navigation | `my.navigateTo` | Page routing |

---

## 📂 Project Structure Details

### Pages

Each page consists of 4 files:
- `.axml` — UI template (markup)
- `.js` — Page logic and lifecycle
- `.acss` — Page-specific styles
- `.json` — Page configuration

**Example**: `pages/listing/`
```
listing/
├── listing.axml      # UI template
├── listing.js        # Page logic
├── listing.acss      # Styles
└── listing.json      # Config
```

### Components

Reusable UI components with the same 4-file structure:

**Example**: `components/search-bar/`
```
search-bar/
├── search-bar.axml   # Component template
├── search-bar.js     # Component logic
├── search-bar.acss   # Component styles
└── search-bar.json   # Component config
```

### Services

Business logic layer that handles:
- Data fetching and caching
- Data transformation
- Error handling
- Local storage operations

### API Layer

HTTP client abstraction:
- **config.js**: API base URL and settings
- **http.client.js**: Request/response wrapper with retry logic
- **property.api.js**: Property-specific API calls

---

## 🎨 UI/UX Features

### Design Principles

- **Responsive**: Adapts to different screen sizes
- **Accessible**: Clear labels and touch targets
- **Performant**: Optimized rendering and caching
- **User-Friendly**: Intuitive navigation and feedback

### Styling

- **Global Styles**: `app.acss` for app-wide styles
- **Component Styles**: Scoped ACSS files for each component
- **Color Scheme**: Blue primary (`#1677FF`), clean whites and grays
- **Typography**: Clear hierarchy with readable font sizes

### Animations

- **Loading Skeletons**: Shimmer effect during data fetch
- **Carousel Transitions**: Smooth image swiping
- **Button Feedback**: Visual feedback on interactions

---

## 🧪 Testing

### Manual Testing Checklist

#### About Page
- [ ] Page loads correctly
- [ ] Content displays properly
- [ ] Navigation works

#### Listing Page
- [ ] Properties load on page open
- [ ] Search filters properties by title
- [ ] Search debounce works (300ms delay)
- [ ] Layout toggle switches between list/grid
- [ ] View mode persists after page reload
- [ ] Pull-to-refresh reloads data
- [ ] Loading skeleton appears during fetch
- [ ] Empty state shows when no results
- [ ] Error state shows on API failure
- [ ] Retry button works on error

#### Detail Page
- [ ] Navigates from listing card
- [ ] Property details display correctly
- [ ] Image carousel swipes smoothly
- [ ] Pagination dots update on swipe
- [ ] Facilities list renders
- [ ] Terms and conditions show
- [ ] Book Now button shows toast
- [ ] Back navigation works

#### Components
- [ ] Search bar accepts input
- [ ] Layout toggle changes icon
- [ ] Property cards display correctly
- [ ] Image carousel handles single image
- [ ] Loading skeleton animates
- [ ] Empty state shows correct message

---

## 🔧 Configuration

### App Configuration (`app.json`)

```json
{
  "pages": [
    "pages/about/about",
    "pages/listing/listing",
    "pages/detail/detail"
  ],
  "window": {
    "defaultTitle": "Property Finder",
    "titleBarColor": "#1677FF",
    "pullRefresh": true
  },
  "tabBar": {
    "items": [
      { "pagePath": "pages/about/about", "name": "About" },
      { "pagePath": "pages/listing/listing", "name": "Listings" }
    ]
  }
}
```

### Constants (`utils/constants.js`)

Key configuration values:
- `API_BASE_URL`: Backend API endpoint
- `SEARCH_DEBOUNCE_MS`: Search delay (300ms)
- `VIEW_MODE`: List/Grid view modes
- `ERROR_MESSAGES`: User-facing error messages

---

## 📡 API Integration

### Endpoints Used

1. **Get Properties**
   - Endpoint: `GET /api/v1/properties`
   - Query: `?search={query}`
   - Used by: Listing page

2. **Get Property Detail**
   - Endpoint: `GET /api/v1/properties/{id}`
   - Used by: Detail page

### Error Handling

- Network errors: Show retry button
- Server errors: Display error message
- Empty results: Show empty state
- Timeout: Configurable retry with exponential backoff

### Caching

- Properties cached in memory
- Cache cleared on pull-to-refresh
- Reduces unnecessary API calls

---

## 🚢 Deployment

### Build for Production

1. **Open DANA Mini Program IDE**
2. **Click Build** → **Upload**
3. **Enter Version Info**:
   - Version: `1.0.0`
   - Description: Initial release
4. **Submit for Review**

### Pre-Deployment Checklist

- [ ] Update API base URL to production
- [ ] Test all features in simulator
- [ ] Test on physical device
- [ ] Verify image loading
- [ ] Check error handling
- [ ] Test pull-to-refresh
- [ ] Verify search functionality
- [ ] Test layout toggle persistence

---

## 🐛 Troubleshooting

### Common Issues

**Properties not loading:**
- Ensure backend is running on `http://localhost:8080`
- Check API configuration in `api/config.js`
- Verify network connectivity

**Search not working:**
- Check debounce delay (should be 300ms)
- Verify search query is trimmed
- Check console for errors

**Layout toggle not persisting:**
- Verify local storage permissions
- Check `storage.service.js` implementation

**Images not loading:**
- Ensure image URLs are accessible
- Check CORS settings on image server
- Verify network connectivity

**Carousel not swiping:**
- Check `swiper` component configuration
- Verify images array is not empty
- Test touch events in simulator

---

## 🔄 Future Enhancements

- [ ] Favorites/wishlist functionality
- [ ] Advanced filters (price, location, facilities)
- [ ] Map view integration
- [ ] User authentication
- [ ] Booking history
- [ ] Push notifications
- [ ] Offline mode with local storage
- [ ] Share property feature
- [ ] Property comparison
- [ ] Virtual tour integration

---

## 📚 Resources

### DANA Mini Program Documentation

- **Official Docs**: [mini-program.dana.id/docs](https://mini-program.dana.id/docs)
- **API Reference**: DANA Mini Program API documentation
- **Component Library**: DANA UI components

### Project Documentation

- **Root README**: [`../README.md`](../README.md)
- **Backend README**: [`../backend/README.md`](../backend/README.md)
- **API Spec**: [`../backend/docs/api-spec.yaml`](../backend/docs/api-spec.yaml)

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
- DANA Mini Program framework expertise
- Component-based architecture
- Modern UI/UX patterns
- Clean, maintainable code
- Production-ready frontend development
