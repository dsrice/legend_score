// apiClient.ts
// This file provides a centralized way to make API requests with the base URL from environment variables
import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';
import { getToken } from './auth';

/**
 * Get the API base URL from environment variables
 * In development, this will use the value from test.env.development or test.env.local
 * In production, this will use the value from test.env.production
 */
export const getApiBaseUrl = (): string => {
  return import.meta.env.VITE_API_BASE_URL || '';
};

// Create an axios instance with the default configuration
const apiClient = axios.create({
  baseURL: getApiBaseUrl(),
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add a request interceptor to include the auth token in all requests
apiClient.interceptors.request.use(
  (config) => {
    // Get the token from localStorage
    const token = getToken();

    // If token exists, add it to the Authorization header
    if (token) {
      config.headers = config.headers || {};
      config.headers.Authorization = `${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

/**
 * Make a GET request to the API
 * @param endpoint The API endpoint (without the base URL)
 * @param params Optional object to be converted to query parameters
 * @param options Additional axios options
 * @returns A promise that resolves to the response data
 */
export const apiGet = async (endpoint: string, params?: Record<string, any>, options?: AxiosRequestConfig): Promise<any> => {
  try {
    const config: AxiosRequestConfig = { ...options };
    if (params) {
      config.params = params;
    }
    const response: AxiosResponse = await apiClient.get(endpoint, config);
    return response.data;
  } catch (error) {
    console.error(`API GET error for ${endpoint}:`, error);
    throw error;
  }
};

/**
 * Make a POST request to the API
 * @param endpoint The API endpoint (without the base URL)
 * @param data The data to send in the request body
 * @param options Additional axios options
 * @returns A promise that resolves to the response data
 */
export const apiPost = async (endpoint: string, data: any, options?: AxiosRequestConfig): Promise<any> => {
  try {
    const response: AxiosResponse = await apiClient.post(endpoint, data, options);
    return response.data;
  } catch (error) {
    console.error(`API POST error for ${endpoint}:`, error);
    throw error;
  }
};

/**
 * Make a PUT request to the API
 * @param endpoint The API endpoint (without the base URL)
 * @param data The data to send in the request body
 * @param options Additional axios options
 * @returns A promise that resolves to the response data
 */
export const apiPut = async (endpoint: string, data: any, options?: AxiosRequestConfig): Promise<any> => {
  try {
    const response: AxiosResponse = await apiClient.put(endpoint, data, options);
    return response.data;
  } catch (error) {
    console.error(`API PUT error for ${endpoint}:`, error);
    throw error;
  }
};

/**
 * Make a DELETE request to the API
 * @param endpoint The API endpoint (without the base URL)
 * @param options Additional axios options
 * @returns A promise that resolves to the response data
 */
export const apiDelete = async (endpoint: string, options?: AxiosRequestConfig): Promise<any> => {
  try {
    const response: AxiosResponse = await apiClient.delete(endpoint, options);
    return response.data;
  } catch (error) {
    console.error(`API DELETE error for ${endpoint}:`, error);
    throw error;
  }
};

// Export the axios instance for more advanced use cases
export default apiClient;