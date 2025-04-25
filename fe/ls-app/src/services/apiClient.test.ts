import { getApiBaseUrl } from './apiClient';

// Mock the process.env object
const originalEnv = process.env;

describe('apiClient', () => {
  beforeEach(() => {
    // Reset process.env before each test
    process.env = { ...originalEnv };
  });

  afterAll(() => {
    // Restore original process.env after all tests
    process.env = originalEnv;
  });

  describe('getApiBaseUrl', () => {
    it('returns empty string when REACT_APP_API_BASE_URL is not set', () => {
      // Ensure the environment variable is not set
      delete process.env.REACT_APP_API_BASE_URL;
      
      expect(getApiBaseUrl()).toBe('');
    });

    it('returns the value of REACT_APP_API_BASE_URL when set', () => {
      // Set the environment variable
      process.env.REACT_APP_API_BASE_URL = 'http://test-api.example.com';
      
      expect(getApiBaseUrl()).toBe('http://test-api.example.com');
    });
  });
});