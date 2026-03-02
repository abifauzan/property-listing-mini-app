/**
 * Creates a debounced version of the given function.
 * @param {Function} fn - Function to debounce
 * @param {number} delay - Delay in milliseconds
 * @returns {Function} Debounced function
 */
function debounce(fn, delay) {
  var timer = null;
  return function () {
    var context = this;
    var args = arguments;
    if (timer) {
      clearTimeout(timer);
    }
    timer = setTimeout(function () {
      fn.apply(context, args);
      timer = null;
    }, delay);
  };
}

module.exports = debounce;
