Component({
  props: {
    images: [],
  },

  data: {
    currentIndex: 0,
  },

  methods: {
    onChange: function (e) {
      this.setData({
        currentIndex: e.detail.current,
      });
    },

    onImageError: function (e) {
      console.warn('[ImageCarousel] Image load failed:', e);
    },

    onPreviewImage: function (e) {
      var index = e.currentTarget.dataset.index;
      var urls = this.props.images || [];
      if (urls.length > 0) {
        my.previewImage({
          current: index,
          urls: urls,
        });
      }
    },
  },
});
