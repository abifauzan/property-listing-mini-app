Component({
  props: {
    mode: 'list',
  },

  methods: {
    onToggle: function () {
      var newMode = this.props.mode === 'list' ? 'grid' : 'list';
      if (this.props.onChange) {
        this.props.onChange(newMode);
      }
    },
  },
});
