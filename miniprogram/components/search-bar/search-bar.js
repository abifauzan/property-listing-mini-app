Component({
  props: {
    value: '',
    placeholder: 'Search properties...',
  },

  methods: {
    onInput: function (e) {
      var value = e.detail.value;
      if (this.props.onChange) {
        this.props.onChange(value);
      }
    },

    onClear: function () {
      if (this.props.onChange) {
        this.props.onChange('');
      }
    },

    onConfirm: function (e) {
      var value = e.detail.value;
      if (this.props.onConfirm) {
        this.props.onConfirm(value);
      }
    },
  },
});
