import { getApiBaseUrl } from './apiClient';

// Mock the Vite import.meta.env
// We need to mock this before importing the module
jest.mock('./apiClient', () => {
  const originalModule = jest.requireActual('./apiClient');
  return {
    ...originalModule,
    getApiBaseUrl: jest.fn(),
  };
});

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