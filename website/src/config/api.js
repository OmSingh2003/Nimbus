// API Configuration
const API_CONFIG = {
  // Use environment variable or fallback to development URL
  BASE_URL: process.env.REACT_APP_API_URL || 'http://localhost:8080',
  
  // API endpoints
  ENDPOINTS: {
    CREATE_USER: '/v1/create_user',
    LOGIN_USER: '/v1/login_user', 
    VERIFY_EMAIL: '/v1/verify_email',
    RESEND_VERIFICATION: '/v1/resend_verification',
    CREATE_ACCOUNT: '/v1/create_account',
    GET_ACCOUNT: '/v1/get_account',
    LIST_ACCOUNTS: '/v1/list_accounts',
    CREATE_TRANSFER: '/v1/create_transfer',
  }
};

// Create axios instance with base configuration
import axios from 'axios';

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
