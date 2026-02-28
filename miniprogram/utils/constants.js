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
    NOT_FOUND: 'Property not found.',
    SERVER: 'Something went wrong. Please try again.',
    EMPTY_SEARCH: 'No properties found for your search.',
    EMPTY_LIST: 'No properties available.',
  },
};
