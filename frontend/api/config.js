const app = getApp();

const config = {
  get baseURL() {
    return (app && app.globalData && app.globalData.apiBaseURL) || 'http://localhost:8080/api/v1';
  },
  timeout: 10000,
  retryAttempts: 2,
  retryDelay: 1000,
};

module.exports = config;
