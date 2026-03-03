module.exports = {
  VIEW_MODE: {
    LIST: 'list',
    GRID: 'grid',
  },

  STORAGE_KEYS: {
    VIEW_MODE: 'property_finder_view_mode',
    LISTING_CACHE: 'property_finder_listing_cache',
    DETAIL_CACHE_PREFIX: 'property_finder_detail_',
  },

  CACHE_TTL: 5 * 60 * 1000, // 5 minutes

  SEARCH_DEBOUNCE_MS: 300,

  ERROR_MESSAGES: {
    NETWORK: 'Unable to connect. Please check your connection.',
    NO_INTERNET: 'No internet connection. Please check your network settings.',
    SLOW_CONNECTION: 'Slow connection detected. This may take a while...',
    TIMEOUT: 'Request timed out. Please try again.',
    NOT_FOUND: 'Property not found.',
    SERVER: 'Something went wrong. Please try again.',
    INVALID_DATA: 'Invalid data received from server.',
    EMPTY_SEARCH: 'No properties found for your search.',
    EMPTY_LIST: 'No properties available.',
    STORAGE_FULL: 'Device storage is full. Some features may not work.',
  },
};
