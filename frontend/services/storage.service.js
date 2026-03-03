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
      storageService._handleStorageError(e);
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
          cachedAt: Date.now(),
        },
      });
    } catch (e) {
      console.error('[StorageService] Failed to cache data:', e);
      storageService._handleStorageError(e);
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

  /**
   * Get cache metadata (timestamp, expiry).
   * @param {string} key - Cache key
   * @returns {Object|null} Metadata or null
   */
  getCacheMetadata: function (key) {
    try {
      var res = my.getStorageSync({ key: key });
      if (res.data && res.data.cachedAt) {
        return {
          cachedAt: res.data.cachedAt,
          expiry: res.data.expiry,
          isExpired: Date.now() >= res.data.expiry,
        };
      }
      return null;
    } catch (e) {
      return null;
    }
  },

  /**
   * Handle storage errors and show appropriate messages.
   * @param {Error} error - Storage error
   */
  _handleStorageError: function (error) {
    var errorMsg = error && error.message ? error.message.toLowerCase() : '';
    
    if (errorMsg.indexOf('quota') !== -1 || errorMsg.indexOf('exceed') !== -1) {
      console.error('[StorageService] Storage quota exceeded');
      my.showToast({
        type: 'fail',
        content: constants.ERROR_MESSAGES.STORAGE_FULL,
        duration: 3000,
      });
    }
  },
};

module.exports = storageService;
