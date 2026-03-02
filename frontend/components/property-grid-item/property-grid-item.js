Component({
  props: {
    property: {},
  },

  methods: {
    onTapCard: function () {
      var property = this.props.property;
      if (property && property.id) {
        my.navigateTo({
          url: '/pages/detail/detail?id=' + property.id,
        });
      }
    },
  },
});
