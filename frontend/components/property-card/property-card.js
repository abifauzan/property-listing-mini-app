Component({
  props: {
    property: {},
  },

  data: {
    imageUrl: '',
    isNavigating: false,
  },

  didMount: function () {
    this.setData({
      imageUrl: this.props.property.imageUrl || '/assets/images/placeholder.png',
    });
  },

  didUpdate: function (prevProps) {
    if (prevProps.property.imageUrl !== this.props.property.imageUrl) {
      this.setData({
        imageUrl: this.props.property.imageUrl || '/assets/images/placeholder.png',
      });
    }
  },

  methods: {
    onImageError: function () {
      console.warn('[PropertyCard] Image load failed, using placeholder');
      this.setData({
        imageUrl: '/assets/images/placeholder.png',
      });
    },

    onTapCard: function () {
      if (this.data.isNavigating) {
        return;
      }

      var property = this.props.property;
      if (property && property.id) {
        this.setData({ isNavigating: true });
        
        var self = this;
        my.navigateTo({
          url: '/pages/detail/detail?id=' + property.id,
          success: function () {
            setTimeout(function () {
              self.setData({ isNavigating: false });
            }, 500);
          },
          fail: function () {
            self.setData({ isNavigating: false });
          },
        });
      }
    },
  },
});
