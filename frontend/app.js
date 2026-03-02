App({
  globalData: {
    apiBaseURL: 'http://localhost:8080/api/v1',
  },

  onLaunch() {
    console.log('Property Finder App launched');
  },

  onShow() {
    console.log('Property Finder App shown');
  },

  onHide() {
    console.log('Property Finder App hidden');
  },
});
