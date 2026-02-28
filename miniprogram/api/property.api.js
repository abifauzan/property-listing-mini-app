var httpClient = require('./http.client');

/**
 * Property API — Data access layer for property endpoints.
 */
var propertyApi = {
  /**
   * Fetch property listings with optional search query.
   * GET /properties?search={query}
   * @param {string} [searchQuery] - Search term to filter by title
   * @returns {Promise<Object>} Response containing { data: { propertyListings: [...] } }
   */
  fetchProperties: function (searchQuery) {
    var params = {};
    if (searchQuery) {
      params.search = searchQuery;
    }
    return httpClient.get('/properties', params);
  },

  /**
   * Fetch property detail by document ID.
   * GET /properties/{id}
   * @param {string} id - Property document ID
   * @returns {Promise<Object>} Response containing { data: { propertyListings: [{...}] } }
   */
  fetchPropertyDetail: function (id) {
    return httpClient.get('/properties/' + id);
  },
};

module.exports = propertyApi;
