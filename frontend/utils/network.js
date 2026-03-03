/**
 * Network utility functions for connectivity detection and status monitoring.
 */
var networkUtil = {
  /**
   * Check if device has network connectivity.
   * @returns {Promise<Object>} { isConnected: boolean, networkType: string }
   */
  checkConnectivity: function () {
    return new Promise(function (resolve) {
      my.getNetworkType({
        success: function (res) {
          var networkType = res.networkType;
          var isConnected = networkType !== 'none' && networkType !== 'unknown';
          resolve({
            isConnected: isConnected,
            networkType: networkType,
          });
        },
        fail: function () {
          resolve({
            isConnected: false,
            networkType: 'unknown',
          });
        },
      });
    });
  },

  /**
   * Check if network type is slow (2G/3G).
   * @param {string} networkType - Network type from my.getNetworkType
   * @returns {boolean}
   */
  isSlowConnection: function (networkType) {
    return networkType === '2g' || networkType === '3g';
  },

  /**
   * Monitor network status changes.
   * @param {Function} callback - Called when network status changes
   */
  onNetworkStatusChange: function (callback) {
    my.onNetworkStatusChange(callback);
  },

  /**
   * Stop monitoring network status changes.
   * @param {Function} callback - The callback to remove
   */
  offNetworkStatusChange: function (callback) {
    if (my.offNetworkStatusChange) {
      my.offNetworkStatusChange(callback);
    }
  },
};

module.exports = networkUtil;
