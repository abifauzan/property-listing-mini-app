var constants = require('../utils/constants');

/**
 * Storage Service — Manages local storage for user preferences and caching.
 */
var storageService = {
  /**
   * Get the persisted view mode preference.
   * @returns {string} 'list' or 'grid'
   */
  getViewMode: function () {
    try {
      var res = my.getStorageSync({ key: constants.STORAGE_KEYS.VIEW_MODE });
      return res.data || constants.VIEW_MODE.LIST;
    } catch (e) {
      return constants.VIEW_MODE.LIST;
    }
  },

  /**
   * Persist the view mode preference.
   * @param {string} mode - 'list' or 'grid'
   */
  setViewMode: function (mode) {
    try {
      my.setStorageSync({
        key: constants.STORAGE_KEYS.VIEW_MODE,
        data: mode,
      });
    } catch (e) {
      console.error('[StorageService] Failed to save view mode:', e);
    }
  },

  /**
   * Get cached data if it exists and hasn't expired.
   * @param {string} key - Cache key
   * @returns {Object|null} Cached data or null
   */
  getCachedData: function (key) {
    try {
      var res = my.getStorageSync({ key: key });
      if (res.data) {
        var cached = res.data;
        if (cached.expiry && Date.now() < cached.expiry) {
          return cached.data;
        }
        my.removeStorageSync({ key: key });
      }
      return null;
    } catch (e) {
      return null;
    }
  },

  /**
   * Store data in cache with a TTL.
   * @param {string} key - Cache key
   * @param {*} data - Data to cache
   * @param {number} [ttl] - Time to live in ms (default: CACHE_TTL from constants)
   */
  setCachedData: function (key, data, ttl) {
    ttl = ttl || constants.CACHE_TTL;
    try {
      my.setStorageSync({
        key: key,
        data: {
          data: data,
          expiry: Date.now() + ttl,
        },
      });
    } catch (e) {
      console.error('[StorageService] Failed to cache data:', e);
    }
  },

  /**
   * Clear all cached property data.
   */
  clearCache: function () {
    try {
      my.removeStorageSync({ key: constants.STORAGE_KEYS.LISTING_CACHE });
    } catch (e) {
      console.error('[StorageService] Failed to clear cache:', e);
    }
  },
};

module.exports = storageService;
