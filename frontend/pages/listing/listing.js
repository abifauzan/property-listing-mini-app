var propertyService = require('../../services/property.service');
var storageService = require('../../services/storage.service');
var debounce = require('../../utils/debounce');
var constants = require('../../utils/constants');

Page({
  data: {
    properties: [],
    allProperties: [],
    searchQuery: '',
    viewMode: constants.VIEW_MODE.LIST,
    loading: true,
    error: null,
    isEmpty: false,
  },

  _debouncedSearch: null,

  onLoad: function () {
    var viewMode = storageService.getViewMode();
    this.setData({ viewMode: viewMode });

    var self = this;
    this._debouncedSearch = debounce(function (query) {
      self._performSearch(query);
    }, constants.SEARCH_DEBOUNCE_MS);

    this._fetchProperties();
  },

  onShow: function () {},

  onPullDownRefresh: function () {
    var self = this;
    propertyService.clearCache();
    this._fetchProperties(true).then(function () {
      my.stopPullDownRefresh();
    }).catch(function () {
      my.stopPullDownRefresh();
    });
  },

  onSearchChange: function (value) {
    this.setData({ searchQuery: value });
    if (!value || value.trim() === '') {
      this.setData({
        properties: this.data.allProperties,
        isEmpty: this.data.allProperties.length === 0,
      });
      return;
    }
    this._debouncedSearch(value);
  },

  onLayoutToggle: function (mode) {
    this.setData({ viewMode: mode });
    storageService.setViewMode(mode);
  },

  onRetry: function () {
    this.setData({ error: null, loading: true });
    this._fetchProperties();
  },

  _fetchProperties: function (forceRefresh) {
    var self = this;
    this.setData({ loading: true, error: null });

    return propertyService.getProperties('', forceRefresh).then(function (properties) {
      self.setData({
        allProperties: properties,
        properties: properties,
        loading: false,
        isEmpty: properties.length === 0,
      });

      if (self.data.searchQuery) {
        self._performSearch(self.data.searchQuery);
      }
    }).catch(function (err) {
      console.error('[ListingPage] Failed to fetch properties:', err);
      self.setData({
        loading: false,
        error: {
          message: err.message || constants.ERROR_MESSAGES.SERVER,
        },
      });
    });
  },

  _performSearch: function (query) {
    var filtered = propertyService.filterProperties(this.data.allProperties, query);
    this.setData({
      properties: filtered,
      isEmpty: filtered.length === 0,
    });
  },
});
