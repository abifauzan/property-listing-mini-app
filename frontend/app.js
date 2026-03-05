const envConfig = require('./config/env');

App({
  globalData: {
    apiBaseURL: envConfig.apiBaseURL,
  },

  onLaunch() {
    console.log('Property Finder App launched');
    console.log('Environment:', envConfig.env);
    console.log('API Base URL:', envConfig.apiBaseURL);
  },

  onShow() {
    console.log('Property Finder App shown');
  },

  onHide() {
    console.log('Property Finder App hidden');
  },

  onError(error) {
    console.error('[Global Error Handler]', error);
    
    var errorMessage = 'An unexpected error occurred';
    if (error && typeof error === 'string') {
      errorMessage = error;
    } else if (error && error.message) {
      errorMessage = error.message;
    }

    my.showToast({
      type: 'fail',
      content: errorMessage,
      duration: 3000,
    });

    return true;
  },
});
