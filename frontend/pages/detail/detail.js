var propertyService = require('../../services/property.service');
var networkUtil = require('../../utils/network');
var sanitizeUtil = require('../../utils/sanitize');
var constants = require('../../utils/constants');

Page({
  data: {
    property: null,
    loading: true,
    error: null,
    isOffline: false,
    offlineMessage: '',
    heroImageUrl: '',
  },

  _propertyId: null,
  _networkStatusHandler: null,

  onLoad: function (query) {
    var sanitizedId = sanitizeUtil.sanitizePropertyId(query.id);
    this._propertyId = sanitizedId;
    
    if (!this._propertyId) {
      this.setData({
        loading: false,
        error: { message: constants.ERROR_MESSAGES.NOT_FOUND },
      });
      return;
    }
    this._setupNetworkMonitoring();
    this._fetchPropertyDetail();
  },

  onShow: function () {
    this._checkNetworkStatus();
  },

  onUnload: function () {
    if (this._networkStatusHandler) {
      networkUtil.offNetworkStatusChange(this._networkStatusHandler);
    }
  },

  onBookNow: function () {
    my.showToast({
      type: 'success',
      content: 'Booking request sent!',
      duration: 2000,
    });
  },

  onRetry: function () {
    this.setData({ error: null, loading: true });
    this._fetchPropertyDetail();
  },

  _fetchPropertyDetail: function () {
    var self = this;
    this.setData({ loading: true, error: null });

    propertyService.getPropertyById(this._propertyId).then(function (property) {
      my.setNavigationBar({ title: property.title });
      self.setData({
        property: property,
        heroImageUrl: property.imageUrl || '/assets/images/placeholder.svg',
        loading: false,
      });
    }).catch(function (err) {
      console.error('[DetailPage] Failed to fetch property detail:', err);
      self.setData({
        loading: false,
        error: {
          message: err.message || constants.ERROR_MESSAGES.SERVER,
        },
      });
    });
  },

  onHeroImageError: function () {
    console.warn('[DetailPage] Hero image load failed, using placeholder');
    this.setData({
      heroImageUrl: '/assets/images/placeholder.svg',
    });
  },

  _setupNetworkMonitoring: function () {
    var self = this;
    this._networkStatusHandler = function (res) {
      var isConnected = res.isConnected;
      self.setData({
        isOffline: !isConnected,
        offlineMessage: isConnected ? '' : constants.ERROR_MESSAGES.NO_INTERNET,
      });
    };
    networkUtil.onNetworkStatusChange(this._networkStatusHandler);
  },

  _checkNetworkStatus: function () {
    var self = this;
    networkUtil.checkConnectivity().then(function (status) {
      self.setData({
        isOffline: !status.isConnected,
        offlineMessage: status.isConnected ? '' : constants.ERROR_MESSAGES.NO_INTERNET,
      });
    });
  },
});
