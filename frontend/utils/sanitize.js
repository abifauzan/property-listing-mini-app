/**
 * Input sanitization utilities to prevent XSS and handle edge cases.
 */
var sanitizeUtil = {
  /**
   * Sanitize search input by removing dangerous characters and limiting length.
   * @param {string} input - Raw user input
   * @param {number} [maxLength] - Maximum allowed length (default: 100)
   * @returns {string} Sanitized input
   */
  sanitizeSearchInput: function (input, maxLength) {
    maxLength = maxLength || 100;

    if (!input || typeof input !== 'string') {
      return '';
    }

    var sanitized = input
      .replace(/[<>]/g, '')
      .replace(/[\x00-\x1F\x7F]/g, '')
      .trim();

    if (sanitized.length > maxLength) {
      sanitized = sanitized.substring(0, maxLength);
    }

    return sanitized;
  },

  /**
   * Validate and sanitize property ID.
   * @param {string} id - Property ID
   * @returns {string|null} Sanitized ID or null if invalid
   */
  sanitizePropertyId: function (id) {
    if (!id || typeof id !== 'string') {
      return null;
    }

    var sanitized = id.trim();
    
    if (sanitized.length === 0 || sanitized.length > 50) {
      return null;
    }

    if (!/^[a-zA-Z0-9_-]+$/.test(sanitized)) {
      return null;
    }

    return sanitized;
  },

  /**
   * Escape HTML special characters to prevent XSS.
   * @param {string} text - Text to escape
   * @returns {string} Escaped text
   */
  escapeHtml: function (text) {
    if (!text || typeof text !== 'string') {
      return '';
    }

    var map = {
      '&': '&amp;',
      '<': '&lt;',
      '>': '&gt;',
      '"': '&quot;',
      "'": '&#x27;',
      '/': '&#x2F;',
    };

    return text.replace(/[&<>"'/]/g, function (char) {
      return map[char];
    });
  },
};

module.exports = sanitizeUtil;
