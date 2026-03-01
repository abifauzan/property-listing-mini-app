# Property Finder — React Native Mobile App

A modern property listing mobile application built with **Expo** and **React Native**, consuming the Go backend API.

---

## 📱 Features

- ✅ **About Screen** — Static information page with app description
- ✅ **Property Listing** — Browse all properties with search and layout toggle
- ✅ **Search** — Real-time debounced search filtering
- ✅ **Layout Toggle** — Switch between list and grid views
- ✅ **Property Detail** — Full property information with gallery, facilities, terms
- ✅ **Pull-to-Refresh** — Refresh property listings
- ✅ **Loading States** — Animated skeleton placeholders
- ✅ **Empty States** — User-friendly messages for no results
- ✅ **Book Now** — Demo booking action with confirmation dialog

---

## 🏗️ Architecture

```
mobile/
├── app/                        # Expo Router file-based routing
│   ├── (tabs)/                 # Bottom tab navigator
│   │   ├── _layout.tsx         # Tab configuration
│   │   ├── index.tsx           # About screen
│   │   └── listing.tsx         # Property listing screen
│   ├── property/
│   │   └── [id].tsx            # Property detail (dynamic route)
│   └── _layout.tsx             # Root layout with React Query provider
├── components/                 # Reusable UI components
│   ├── PropertyCard.tsx        # Vertical list card
│   ├── PropertyGridItem.tsx    # Grid view card
│   ├── SearchBar.tsx           # Search input with debounce
│   ├── LayoutToggle.tsx        # List/Grid toggle button
│   ├── ImageCarousel.tsx       # Gallery carousel
│   ├── EmptyState.tsx          # No results state
│   └── LoadingSkeleton.tsx     # Skeleton placeholders
├── hooks/                      # Custom React hooks
│   ├── useProperties.ts        # Fetch property listings
│   └── usePropertyDetail.ts    # Fetch single property
├── services/                   # API layer
│   ├── api.client.ts           # Axios instance with retry logic
│   └── property.api.ts         # Typed API functions
├── types/                      # TypeScript interfaces
│   └── property.ts             # Property domain models
├── utils/                      # Pure utility functions
│   ├── format.ts               # formatPrice, formatDate
│   ├── debounce.ts             # Debounce utility
│   └── constants.ts            # API URLs, colors, query keys
└── assets/                     # Static assets
```

---

## 🛠️ Technology Stack

| Layer | Technology | Purpose |
|-------|-----------|---------|
| **Framework** | Expo SDK 55 | Fast development, OTA updates |
| **Navigation** | Expo Router v4 | File-based routing, deep linking |
| **Styling** | NativeWind v4 | TailwindCSS for React Native |
| **Data Fetching** | TanStack Query v5 | Caching, stale-while-revalidate |
| **HTTP Client** | Axios | Retry logic, interceptors |
| **Language** | TypeScript (strict) | Type safety |
| **Icons** | @expo/vector-icons | Ionicons |
| **Images** | expo-image | Progressive loading, caching |

---

## 📋 Prerequisites

- **Node.js** 18.x or higher
- **npm** or **yarn**
- **Expo CLI** (optional, npx works)
- **Expo Go** app on iOS/Android device (for testing)
- **Backend API** running at `http://localhost:8080`

---

## 🚀 Getting Started

### 1. Install Dependencies

```bash
cd mobile
npm install
```

### 2. Start the Backend API

Make sure the Go backend is running:

```bash
cd ../backend
go run cmd/api/main.go
```

The API should be accessible at `http://localhost:8080/api/v1`.

### 3. Start the Expo Development Server

```bash
npm start
```

### 4. Run on Device/Simulator

- **iOS Simulator**: Press `i` in the terminal
- **Android Emulator**: Press `a` in the terminal
- **Physical Device**: Scan the QR code with Expo Go app

---

## 📱 Running on Physical Device

### iOS (via Expo Go)

1. Install **Expo Go** from the App Store
2. Scan the QR code from the terminal
3. **Important**: Your phone and computer must be on the same network
4. If using `localhost` API, you may need to:
   - Update `API_BASE_URL` in `utils/constants.ts` to your computer's local IP
   - Example: `http://192.168.1.100:8080`

### Android (via Expo Go)

1. Install **Expo Go** from Google Play
2. Scan the QR code from the terminal
3. Same network requirements as iOS

---

## 🔧 Configuration

### API Base URL

Update the backend API URL in `utils/constants.ts`:

```typescript
export const API_BASE_URL = 'http://localhost:8080';
// For physical device testing, use your computer's IP:
// export const API_BASE_URL = 'http://192.168.1.100:8080';
```

### Colors

Primary brand color is defined in `utils/constants.ts` and `tailwind.config.js`:

```typescript
PRIMARY: '#0060FF'  // DANA blue
```

---

## 📂 Key Files

| File | Purpose |
|------|---------|
| `app/_layout.tsx` | Root layout, React Query provider |
| `app/(tabs)/_layout.tsx` | Tab bar configuration |
| `app/(tabs)/index.tsx` | About screen |
| `app/(tabs)/listing.tsx` | Property listing with search & toggle |
| `app/property/[id].tsx` | Property detail screen |
| `services/api.client.ts` | Axios instance with retry & timeout |
| `hooks/useProperties.ts` | React Query hook for listings |
| `utils/constants.ts` | API URL, colors, query keys |

---

## 🎨 Design Decisions

### Why Expo Router?

- **File-based routing** — Convention over configuration
- **Type-safe navigation** — Auto-generated typed routes
- **Deep linking** — Built-in support for universal links
- **Modern standard** — Aligns with Expo's recommended approach

### Why NativeWind?

- **Utility-first** — Faster iteration than StyleSheet
- **Consistent design tokens** — Shared spacing, colors
- **Responsive utilities** — Easy breakpoint handling
- **Familiar** — Same API as TailwindCSS

### Why TanStack Query?

- **Automatic caching** — Reduces unnecessary API calls
- **Background refetches** — Keeps data fresh
- **Pull-to-refresh** — Built-in integration
- **Loading/error states** — Declarative state management

---

## 🧪 Testing

### Manual Testing Checklist

- [ ] About screen displays correctly
- [ ] Property listing loads all items
- [ ] Search filters properties by title
- [ ] Layout toggle switches between list/grid
- [ ] Pull-to-refresh reloads data
- [ ] Tapping a property navigates to detail
- [ ] Detail screen shows all property info
- [ ] Image carousel works with pagination dots
- [ ] "Book Now" button shows confirmation dialog
- [ ] Empty state appears when search has no results
- [ ] Loading skeleton appears during data fetch

---

## 🐛 Troubleshooting

### "Network request failed"

- Ensure backend is running at `http://localhost:8080`
- For physical devices, update `API_BASE_URL` to your computer's local IP
- Check that device and computer are on the same network

### TypeScript errors

- Run `npm install` to ensure all dependencies are installed
- Restart the TypeScript server in your IDE

### Metro bundler issues

```bash
npm start -- --clear
```

---

## 📦 Build for Production

### iOS

```bash
npx eas build --platform ios
```

### Android

```bash
npx eas build --platform android
```

**Note**: Requires an Expo account and EAS CLI setup.

---

## 🔄 API Contract

The app consumes the following backend endpoints:

| Endpoint | Method | Query Params | Response |
|----------|--------|--------------|----------|
| `/api/v1/properties` | GET | `?search={query}` | `{ data: { propertyListings: PropertyListing[] } }` |
| `/api/v1/properties/{id}` | GET | — | `{ data: { propertyListings: Property[] } }` |

---

## 📸 Screenshots

*(Add screenshots here after running the app)*

---

## 🚧 Future Enhancements

- [ ] Add property favorites (local storage)
- [ ] Implement property filtering (price range, facilities)
- [ ] Add map view for property locations
- [ ] Implement authentication
- [ ] Add booking history
- [ ] Push notifications for new properties
- [ ] Offline mode with local caching

---

## 📄 License

MIT

---

## 👤 Author

Abi Fauzan  
Email: abifauzan234@gmail.com
