/**
 * Format a numeric price string to Indonesian Rupiah format.
 * e.g., "750000" → "Rp 750.000"
 * @param {string|number} price - Raw price value
 * @returns {string} Formatted price string
 */
function formatPrice(price) {
  if (price === undefined || price === null || price === '') {
    return 'Rp 0';
  }
  var num = typeof price === 'string' ? parseInt(price, 10) : price;
  if (isNaN(num)) {
    return 'Rp 0';
  }
  var formatted = num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
  return 'Rp ' + formatted;
}

/**
 * Format ISO date string to readable format.
 * e.g., "2025-01-31T08:18:56.778Z" → "31 Jan 2025"
 * @param {string} dateStr - ISO date string
 * @returns {string} Formatted date string
 */
function formatDate(dateStr) {
  if (!dateStr) return '';
  var months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  var d = new Date(dateStr);
  if (isNaN(d.getTime())) return '';
  return d.getDate() + ' ' + months[d.getMonth()] + ' ' + d.getFullYear();
}

module.exports = {
  formatPrice: formatPrice,
  formatDate: formatDate,
};
