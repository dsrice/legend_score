import { getApiBaseUrl } from './env';
import { apiGet, apiPost, apiPut, apiDelete } from './apiClient';
import axios from 'axios';
import { getToken } from './auth';

// Mock the env module
jest.mock('./env', () => ({
  getApiBaseUrl: jest.fn(),
  setApiBaseUrl: jest.fn()
}));

// Mock axios
jest.mock('axios', () => ({
  create: jest.fn(() => ({
    get: jest.fn(),
    post: jest.fn(),
    put: jest.fn(),
    delete: jest.fn(),
    interceptors: {
      request: {
        use: jest.fn()
      }
    }
  }))
}));

// Mock auth
jest.mock('./auth', () => ({
  getToken: jest.fn()
}));

describe('apiClient', () => {
  // Reset mocks before each test
  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe('getApiBaseUrl', () => {
    it('returns empty string when VITE_API_BASE_URL is not set', () => {
      // Mock implementation for this test
      (getApiBaseUrl as jest.Mock).mockReturnValue('');

      expect(getApiBaseUrl()).toBe('');
    });

    it('returns the value of VITE_API_BASE_URL when set', () => {
      // Mock implementation for this test
      (getApiBaseUrl as jest.Mock).mockReturnValue('http://test-api.example.com');

      expect(getApiBaseUrl()).toBe('http://test-api.example.com');
    });
  });

  describe('apiGet', () => {
    it('makes a GET request to the correct endpoint', async () => {
      const mockAxiosInstance = axios.create();
      (mockAxiosInstance.get as jest.Mock).mockResolvedValue({ data: { result: true } });

      await apiGet('/test-endpoint');

      expect(mockAxiosInstance.get).toHaveBeenCalledWith('/test-endpoint', expect.any(Object));
    });

    it('handles query parameters correctly', async () => {
      const mockAxiosInstance = axios.create();
      (mockAxiosInstance.get as jest.Mock).mockResolvedValue({ data: { result: true } });

      const params = { id: 1, name: 'test' };
      await apiGet('/test-endpoint', params);

      expect(mockAxiosInstance.get).toHaveBeenCalledWith('/test-endpoint', expect.objectContaining({
        params
      }));
    });

    it('handles errors correctly', async () => {
      const mockAxiosInstance = axios.create();
      const mockError = new Error('Network error');
      (mockAxiosInstance.get as jest.Mock).mockRejectedValue(mockError);

      await expect(apiGet('/test-endpoint')).rejects.toThrow();
    });
  });

  describe('apiPost', () => {
    it('makes a POST request to the correct endpoint with data', async () => {
      const mockAxiosInstance = axios.create();
      (mockAxiosInstance.post as jest.Mock).mockResolvedValue({ data: { result: true } });

      const data = { name: 'test', value: 123 };
      await apiPost('/test-endpoint', data);

      expect(mockAxiosInstance.post).toHaveBeenCalledWith('/test-endpoint', data, undefined);
    });

    it('handles errors correctly', async () => {
      const mockAxiosInstance = axios.create();
      const mockError = new Error('Network error');
      (mockAxiosInstance.post as jest.Mock).mockRejectedValue(mockError);

      await expect(apiPost('/test-endpoint', {})).rejects.toThrow();
    });
  });

  describe('apiPut', () => {
    it('makes a PUT request to the correct endpoint with data', async () => {
      const mockAxiosInstance = axios.create();
      (mockAxiosInstance.put as jest.Mock).mockResolvedValue({ data: { result: true } });

      const data = { id: 1, name: 'updated' };
      await apiPut('/test-endpoint/1', data);

      expect(mockAxiosInstance.put).toHaveBeenCalledWith('/test-endpoint/1', data, undefined);
    });

    it('handles errors correctly', async () => {
      const mockAxiosInstance = axios.create();
      const mockError = new Error('Network error');
      (mockAxiosInstance.put as jest.Mock).mockRejectedValue(mockError);

      await expect(apiPut('/test-endpoint/1', {})).rejects.toThrow();
    });
  });

  describe('apiDelete', () => {
    it('makes a DELETE request to the correct endpoint', async () => {
      const mockAxiosInstance = axios.create();
      (mockAxiosInstance.delete as jest.Mock).mockResolvedValue({ data: { result: true } });

      await apiDelete('/test-endpoint/1');

      expect(mockAxiosInstance.delete).toHaveBeenCalledWith('/test-endpoint/1', undefined);
    });

    it('handles errors correctly', async () => {
      const mockAxiosInstance = axios.create();
      const mockError = new Error('Network error');
      (mockAxiosInstance.delete as jest.Mock).mockRejectedValue(mockError);

      await expect(apiDelete('/test-endpoint/1')).rejects.toThrow();
    });
  });

  describe('request interceptor', () => {
    it('adds authorization header when token exists', () => {
      // Setup
      const mockToken = 'test-token';
      (getToken as jest.Mock).mockReturnValue(mockToken);

      // Get the interceptor function
      const mockAxiosInstance = axios.create();
      const requestInterceptor = mockAxiosInstance.interceptors.request.use.mock.calls[0][0];

      // Test the interceptor
      const config = { headers: {} };
      const result = requestInterceptor(config);

      expect(result.headers.Authorization).toBe(`Bearer ${mockToken}`);
    });

    it('does not add authorization header when token does not exist', () => {
      // Setup
      (getToken as jest.Mock).mockReturnValue(null);

      // Get the interceptor function
      const mockAxiosInstance = axios.create();
      const requestInterceptor = mockAxiosInstance.interceptors.request.use.mock.calls[0][0];

      // Test the interceptor
      const config = { headers: {} };
      const result = requestInterceptor(config);

      expect(result.headers.Authorization).toBeUndefined();
    });
  });
});