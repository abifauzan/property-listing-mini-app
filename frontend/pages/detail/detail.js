var propertyService = require('../../services/property.service');
var constants = require('../../utils/constants');

Page({
  data: {
    property: null,
    loading: true,
    error: null,
  },

  _propertyId: null,

  onLoad: function (query) {
    this._propertyId = query.id;
    if (!this._propertyId) {
      this.setData({
        loading: false,
        error: { message: constants.ERROR_MESSAGES.NOT_FOUND },
      });
      return;
    }
    this._fetchPropertyDetail();
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
});
