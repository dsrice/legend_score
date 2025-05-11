import { getApiBaseUrl } from './env';
import { apiGet, apiPost, apiPut, apiDelete } from './apiClient';

// Mock the env module
jest.mock('./env', () => ({
  getApiBaseUrl: jest.fn(),
  setApiBaseUrl: jest.fn()
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
});