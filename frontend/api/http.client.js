const config = require('./config');

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
          if (res.status >= 200 && res.status < 300) {
            resolve(res.data);
          } else {
            var error = httpClient._parseError(res);
            reject(error);
          }
        },
        fail: function (err) {
          if (attempt < config.retryAttempts) {
            setTimeout(function () {
              httpClient._request(method, path, params, attempt + 1)
                .then(resolve)
                .catch(reject);
            }, config.retryDelay * (attempt + 1));
          } else {
            reject({
              message: err.errorMessage || 'Network request failed',
              code: 'NETWORK_ERROR',
            });
          }
        },
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
   * Parse error from response.
   * @param {Object} res - Raw response
   * @returns {Object} Structured error
   */
  _parseError(res) {
    var data = res.data || {};
    var errorData = data.error || {};
    return {
      message: errorData.message || 'An unexpected error occurred',
      code: errorData.code || 'UNKNOWN_ERROR',
      status: res.status,
    };
  },
};

module.exports = httpClient;
