import axios from 'axios';

// API Configuration
const API_CONFIG = {
  // Use environment variable or fallback to Render production URL
  BASE_URL: process.env.REACT_APP_API_URL || 'https://nimbus-91j2.onrender.com',
  
  // API endpoints (gRPC Gateway paths)
  ENDPOINTS: {
    CREATE_USER: '/v1/create_user',
    LOGIN_USER: '/v1/login_user', 
    VERIFY_EMAIL: '/v1/verify_email',
    RESEND_VERIFICATION: '/v1/resend_verification',
    CREATE_ACCOUNT: '/v1/accounts',
    GET_ACCOUNT: '/v1/accounts',
    LIST_ACCOUNTS: '/v1/accounts',
    CREATE_TRANSFER: '/v1/transfers',
    RENEW_TOKEN: '/v1/renew_access_token',
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

// Add response interceptor to handle auth errors
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    // Check if the error is related to authentication
    if (error.response?.status === 401 || 
        (error.response?.status === 500 && 
         error.response?.data?.message?.includes('missing authorization header'))) {
      
      // Clear invalid token and redirect to login
      localStorage.removeItem('token');
      localStorage.removeItem('username');
      
      // Redirect to login page if not already there
      if (window.location.pathname !== '/login' && window.location.pathname !== '/create-user') {
        window.location.href = '/login';
      }
    }
    return Promise.reject(error);
  }
);

export { API_CONFIG, apiClient };
export default apiClient;
