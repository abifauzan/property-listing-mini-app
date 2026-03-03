const config = require('./config');
const networkUtil = require('../utils/network');
const constants = require('../utils/constants');

/**
 * HTTP client wrapper around my.request with retry, timeout, and error handling.
 */
const httpClient = {
  /**
   * Perform an HTTP GET request.
   * @param {string} path - API path (e.g., '/properties')
   * @param {Object} [params] - Query parameters
   * @returns {Promise<Object>} Parsed response data
   */
  get(path, params) {
    return this._request('GET', path, params);
  },

  /**
   * Internal request method with retry logic.
   * @param {string} method - HTTP method
   * @param {string} path - API path
   * @param {Object} [params] - Query parameters
   * @param {number} [attempt] - Current retry attempt
   * @returns {Promise<Object>}
   */
  _request(method, path, params, attempt) {
    attempt = attempt || 0;
    const url = this._buildURL(path, params);

    return networkUtil.checkConnectivity().then(function (networkStatus) {
      if (!networkStatus.isConnected) {
        return Promise.reject({
          message: constants.ERROR_MESSAGES.NO_INTERNET,
          code: 'NO_INTERNET',
        });
      }

      if (networkUtil.isSlowConnection(networkStatus.networkType)) {
        console.warn('[HTTP Client] Slow connection detected:', networkStatus.networkType);
      }

      return new Promise(function (resolve, reject) {
        my.request({
          url: url,
          method: method,
          headers: {
            'Content-Type': 'application/json',
          },
          timeout: config.timeout,
          dataType: 'json',
          success: function (res) {
            try {
              if (res.status >= 200 && res.status < 300) {
                var data = httpClient._safeParseResponse(res.data);
                resolve(data);
              } else {
                var error = httpClient._parseError(res);
                reject(error);
              }
            } catch (parseError) {
              console.error('[HTTP Client] Response parsing error:', parseError);
              reject({
                message: constants.ERROR_MESSAGES.INVALID_DATA,
                code: 'PARSE_ERROR',
              });
            }
          },
          fail: function (err) {
            var isTimeout = err.error === 13 || (err.errorMessage && err.errorMessage.indexOf('timeout') !== -1);
            
            if (isTimeout) {
              reject({
                message: constants.ERROR_MESSAGES.TIMEOUT,
                code: 'TIMEOUT',
              });
              return;
            }

            if (attempt < config.retryAttempts) {
              console.log('[HTTP Client] Retrying request, attempt:', attempt + 1);
              setTimeout(function () {
                httpClient._request(method, path, params, attempt + 1)
                  .then(resolve)
                  .catch(reject);
              }, config.retryDelay * (attempt + 1));
            } else {
              reject({
                message: err.errorMessage || constants.ERROR_MESSAGES.NETWORK,
                code: 'NETWORK_ERROR',
              });
            }
          },
        });
      });
    });
  },

  /**
   * Build full URL with query parameters.
   * @param {string} path - API path
   * @param {Object} [params] - Query parameters
   * @returns {string}
   */
  _buildURL(path, params) {
    var url = config.baseURL + path;
    if (params) {
      var queryParts = [];
      var keys = Object.keys(params);
      for (var i = 0; i < keys.length; i++) {
        var key = keys[i];
        var value = params[key];
        if (value !== undefined && value !== null && value !== '') {
          queryParts.push(encodeURIComponent(key) + '=' + encodeURIComponent(value));
        }
      }
      if (queryParts.length > 0) {
        url += '?' + queryParts.join('&');
      }
    }
    return url;
  },

  /**
   * Safely parse response data, handling malformed JSON.
   * @param {*} data - Response data
   * @returns {*} Parsed data
   */
  _safeParseResponse(data) {
    if (typeof data === 'string') {
      try {
        return JSON.parse(data);
      } catch (e) {
        console.error('[HTTP Client] Failed to parse JSON string:', e);
        throw new Error('Invalid JSON response');
      }
    }
    return data;
  },

  /**
   * Parse error from response.
   * @param {Object} res - Raw response
   * @returns {Object} Structured error
   */
  _parseError(res) {
    var data = res.data || {};
    var errorData = data.error || {};
    var statusMessages = {
      400: 'Bad request',
      401: 'Unauthorized',
      403: 'Forbidden',
      404: 'Not found',
      500: 'Server error',
      502: 'Bad gateway',
      503: 'Service unavailable',
    };

    var message = errorData.message || statusMessages[res.status] || constants.ERROR_MESSAGES.SERVER;
    
    return {
      message: message,
      code: errorData.code || 'HTTP_' + res.status,
      status: res.status,
    };
  },
};

module.exports = httpClient;
