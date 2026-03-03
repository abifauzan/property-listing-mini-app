var propertyApi = require('../api/property.api');
var storageService = require('./storage.service');
var formatUtil = require('../utils/format');
var constants = require('../utils/constants');

/**
 * Property Service — Business logic layer for property data.
 * Handles fetching, caching, transformation, and filtering.
 */
var propertyService = {
  /**
   * Fetch property listings. Uses cache when available.
   * @param {string} [searchQuery] - Optional search term
   * @param {boolean} [forceRefresh] - Skip cache if true
   * @returns {Promise<Array>} Array of formatted property items
   */
  getProperties: function (searchQuery, forceRefresh) {
    var cacheKey = constants.STORAGE_KEYS.LISTING_CACHE;

    if (!forceRefresh && !searchQuery) {
      var cached = storageService.getCachedData(cacheKey);
      if (cached) {
        return Promise.resolve(cached);
      }
    }

    return propertyApi.fetchProperties(searchQuery).then(function (response) {
      var listings = propertyService._extractListings(response);
      var formatted = propertyService._formatListings(listings);

      if (!searchQuery) {
        storageService.setCachedData(cacheKey, formatted);
      }

      return formatted;
    });
  },

  /**
   * Fetch single property detail by ID.
   * @param {string} id - Property document ID
   * @returns {Promise<Object>} Formatted property detail
   */
  getPropertyById: function (id) {
    var cacheKey = constants.STORAGE_KEYS.DETAIL_CACHE_PREFIX + id;

    var cached = storageService.getCachedData(cacheKey);
    if (cached) {
      return Promise.resolve(cached);
    }

    return propertyApi.fetchPropertyDetail(id).then(function (response) {
      var listings = propertyService._extractListings(response);
      if (!listings || listings.length === 0) {
        return Promise.reject({
          message: constants.ERROR_MESSAGES.NOT_FOUND,
          code: 'NOT_FOUND',
        });
      }
      var detail = propertyService._formatDetail(listings[0]);
      storageService.setCachedData(cacheKey, detail);
      return detail;
    });
  },

  /**
   * Filter properties locally by title match.
   * @param {Array} properties - Full property list
   * @param {string} query - Search query
   * @returns {Array} Filtered properties
   */
  filterProperties: function (properties, query) {
    if (!query || !query.trim()) {
      return properties;
    }
    var lowerQuery = query.toLowerCase().trim();
    return properties.filter(function (item) {
      return item.title && item.title.toLowerCase().indexOf(lowerQuery) !== -1;
    });
  },

  /**
   * Clear all property caches. Used for pull-to-refresh.
   */
  clearCache: function () {
    storageService.clearCache();
  },

  // ---- Private helpers ----

  /**
   * Extract property listings array from API response.
   * @param {Object} response - Raw API response
   * @returns {Array}
   */
  _extractListings: function (response) {
    if (response && response.data && response.data.propertyListings) {
      return response.data.propertyListings;
    }
    return [];
  },

  /**
   * Validate if a property item has required fields.
   * @param {Object} item - Property item to validate
   * @returns {boolean}
   */
  _isValidProperty: function (item) {
    if (!item || typeof item !== 'object') {
      return false;
    }
    
    var hasId = item.documentId && typeof item.documentId === 'string';
    var hasTitle = item.Title && typeof item.Title === 'string' && item.Title.trim() !== '';
    var hasPrice = (typeof item.Price === 'number' || typeof item.Price === 'string') && 
                   !isNaN(parseFloat(item.Price)) && 
                   parseFloat(item.Price) >= 0;
    
    return hasId && hasTitle && hasPrice;
  },

  /**
   * Format listing items for display.
   * @param {Array} listings - Raw property listing array
   * @returns {Array} Formatted listings
   */
  _formatListings: function (listings) {
    if (!Array.isArray(listings)) {
      console.warn('[PropertyService] Invalid listings array:', listings);
      return [];
    }

    return listings
      .filter(function (item) {
        var isValid = propertyService._isValidProperty(item);
        if (!isValid) {
          console.warn('[PropertyService] Skipping invalid property:', item);
        }
        return isValid;
      })
      .map(function (item) {
        try {
          var numericPrice = typeof item.Price === 'string' ? parseFloat(item.Price) : item.Price;
          return {
            id: item.documentId,
            title: item.Title || '',
            price: formatUtil.formatPrice(numericPrice),
            rawPrice: numericPrice,
            imageUrl: (item.Banner && item.Banner.url) || '',
            createdAt: formatUtil.formatDate(item.createdAt),
          };
        } catch (err) {
          console.error('[PropertyService] Error formatting property:', err, item);
          return null;
        }
      })
      .filter(function (item) {
        return item !== null;
      });
  },

  /**
   * Format a single property detail for display.
   * @param {Object} item - Raw property object
   * @returns {Object} Formatted property detail
   */
  _formatDetail: function (item) {
    if (!propertyService._isValidProperty(item)) {
      throw new Error('Invalid property data');
    }

    try {
      var images = [];
      if (Array.isArray(item.Images) && item.Images.length > 0) {
        images = item.Images
          .filter(function (img) {
            return img && img.url && typeof img.url === 'string';
          })
          .map(function (img) {
            return img.url;
          });
      }

      var facilities = [];
      if (Array.isArray(item.Facilities)) {
        facilities = item.Facilities.filter(function (f) {
          return f && typeof f === 'string' && f.trim() !== '';
        });
      }

      return {
        id: item.documentId,
        title: item.Title || '',
        price: formatUtil.formatPrice(item.Price),
        rawPrice: item.Price,
        imageUrl: (item.Banner && item.Banner.url) || '',
        description: item.Description || '',
        images: images,
        facilities: facilities,
        terms: item.Terms || '',
        conditions: item.Conditions || '',
        createdAt: formatUtil.formatDate(item.createdAt),
      };
    } catch (err) {
      console.error('[PropertyService] Error formatting property detail:', err);
      throw err;
    }
  },
};

module.exports = propertyService;
