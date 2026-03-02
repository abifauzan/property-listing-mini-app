Component({
  props: {
    type: 'list',
    count: 3,
  },

  data: {
    items: [],
  },

  didMount: function () {
    var count = this.props.count || 3;
    var items = [];
    for (var i = 0; i < count; i++) {
      items.push(i);
    }
    this.setData({ items: items });
  },
});
