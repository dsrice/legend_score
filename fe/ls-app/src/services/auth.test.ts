import { login, storeToken, getToken, removeToken, isAuthenticated } from './auth';
import * as apiClient from './apiClient';

// Mock the apiClient
jest.mock('./apiClient', () => ({
  apiPost: jest.fn()
}));

// Mock localStorage
const localStorageMock = (() => {
  let store: Record<string, string> = {};
  return {
    getItem: jest.fn((key: string) => store[key] || null),
    setItem: jest.fn((key: string, value: string) => {
      store[key] = value;
    }),
    removeItem: jest.fn((key: string) => {
      delete store[key];
    }),
    clear: jest.fn(() => {
      store = {};
    })
  };
})();

Object.defineProperty(window, 'localStorage', { value: localStorageMock });

describe('Auth Service', () => {
  beforeEach(() => {
    jest.clearAllMocks();
    localStorageMock.clear();
  });

  describe('login', () => {
    test('calls apiPost with correct parameters', async () => {
      const credentials = {
        login_id: 'testuser',
        password: 'password123'
      };
      
      const mockResponse = {
        result: true,
        token: 'fake-token'
      };
      
      (apiClient.apiPost as jest.Mock).mockResolvedValue(mockResponse);
      
      const result = await login(credentials);
      
      expect(apiClient.apiPost).toHaveBeenCalledWith('/login', credentials);
      expect(result).toEqual(mockResponse);
    });

    test('throws error when API call fails', async () => {
      const credentials = {
        login_id: 'testuser',
        password: 'password123'
      };
      
      const mockError = new Error('API error');
      (apiClient.apiPost as jest.Mock).mockRejectedValue(mockError);
      
      await expect(login(credentials)).rejects.toThrow('API error');
    });
  });

  describe('storeToken', () => {
    test('stores token in localStorage', () => {
      storeToken('fake-token');
      
      expect(localStorageMock.setItem).toHaveBeenCalledWith('auth_token', 'fake-token');
    });
  });

  describe('getToken', () => {
    test('retrieves token from localStorage', () => {
      localStorageMock.getItem.mockReturnValue('fake-token');
      
      const token = getToken();
      
      expect(localStorageMock.getItem).toHaveBeenCalledWith('auth_token');
      expect(token).toBe('fake-token');
    });

    test('returns null when token is not found', () => {
      localStorageMock.getItem.mockReturnValue(null);
      
      const token = getToken();
      
      expect(token).toBeNull();
    });
  });

  describe('removeToken', () => {
    test('removes token from localStorage', () => {
      removeToken();
      
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('auth_token');
    });
  });

  describe('isAuthenticated', () => {
    test('returns true when token exists', () => {
      localStorageMock.getItem.mockReturnValue('fake-token');
      
      const result = isAuthenticated();
      
      expect(result).toBe(true);
    });

    test('returns false when token does not exist', () => {
      localStorageMock.getItem.mockReturnValue(null);
      
      const result = isAuthenticated();
      
      expect(result).toBe(false);
    });
  });
});