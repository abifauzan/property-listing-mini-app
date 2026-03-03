var sanitizeUtil = require('../../utils/sanitize');

Component({
  props: {
    value: '',
    placeholder: 'Search properties...',
  },

  methods: {
    onInput: function (e) {
      var value = e.detail.value;
      var sanitized = sanitizeUtil.sanitizeSearchInput(value, 100);
      if (this.props.onChange) {
        this.props.onChange(sanitized);
      }
    },

    onClear: function () {
      if (this.props.onChange) {
        this.props.onChange('');
      }
    },

    onConfirm: function (e) {
      var value = e.detail.value;
      var sanitized = sanitizeUtil.sanitizeSearchInput(value, 100);
      if (this.props.onConfirm) {
        this.props.onConfirm(sanitized);
      }
    },
  },
});
