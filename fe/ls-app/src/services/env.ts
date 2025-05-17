// env.ts - Environment variables
// This file is designed to be easily mocked in tests

// Default API base URL
let apiBaseUrl = '';

// In a browser environment, try to get the API base URL from Vite's import.meta.env
// This code will only run in the browser, not in Jest
if (typeof import.meta !== 'undefined') {
  try {
    // Access the environment variable directly from import.meta.env
    if (import.meta.env.VITE_API_BASE_URL) {
      apiBaseUrl = import.meta.env.VITE_API_BASE_URL;
    }
  } catch (error) {
    console.warn('Failed to access Vite environment variables:', error);
  }
}

// Export the API base URL
export const getApiBaseUrl = (): string => {
  return apiBaseUrl;
};

// Allow setting the API base URL (useful for tests)
export const setApiBaseUrl = (url: string): void => {
  apiBaseUrl = url;
};