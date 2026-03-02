import axios, { AxiosError, AxiosInstance, InternalAxiosRequestConfig } from 'axios';
import { API_BASE_URL } from '@/utils/constants';

const MAX_RETRIES = 3;
const RETRY_DELAY = 1000;

const apiClient: AxiosInstance = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

apiClient.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

apiClient.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const config = error.config as InternalAxiosRequestConfig & { _retryCount?: number };
    
    if (!config) {
      return Promise.reject(error);
    }

    config._retryCount = config._retryCount || 0;

    if (config._retryCount >= MAX_RETRIES) {
      return Promise.reject(error);
    }

    if (error.response?.status && error.response.status >= 500) {
      config._retryCount += 1;
      
      await new Promise((resolve) => setTimeout(resolve, RETRY_DELAY * config._retryCount!));
      
      return apiClient(config);
    }

    return Promise.reject(error);
  }
);

export default apiClient;
