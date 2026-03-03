Component({
  props: {
    message: 'No results found',
    showRetry: false,
  },

  methods: {
    onRetry: function () {
      if (this.props.onRetry && typeof this.props.onRetry === 'function') {
        this.props.onRetry();
      }
    },
  },
});
