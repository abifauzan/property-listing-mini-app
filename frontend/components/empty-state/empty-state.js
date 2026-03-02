Component({
  props: {
    message: 'No results found',
    showRetry: false,
  },

  methods: {
    onRetry: function () {
      if (this.props.onRetry) {
        this.props.onRetry();
      }
    },
  },
});
