import axios from 'axios';

// API Configuration
const API_CONFIG = {
  // Use environment variable or fallback to development URL
  BASE_URL: process.env.REACT_APP_API_URL || 'http://localhost:8080',
  
  // API endpoints
  ENDPOINTS: {
    CREATE_USER: '/users',
    LOGIN_USER: '/users/login', 
    VERIFY_EMAIL: '/verify_email',
    RESEND_VERIFICATION: '/resend_verification',
    CREATE_ACCOUNT: '/accounts',
    GET_ACCOUNT: '/accounts',
    LIST_ACCOUNTS: '/accounts',
    CREATE_TRANSFER: '/transfers',
    RENEW_TOKEN: '/token.renew_access',
  }
};

// Create axios instance with base configuration

const apiClient = axios.create({
  baseURL: API_CONFIG.BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add request interceptor to include auth token
apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export { API_CONFIG, apiClient };
export default apiClient;
