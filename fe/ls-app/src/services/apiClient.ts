// apiClient.ts
// This file provides a centralized way to make API requests with the base URL from environment variables

/**
 * Get the API base URL from environment variables
 * In development, this will use the value from .env.development or .env.local
 * In production, this will use the value from .env.production
 */
export const getApiBaseUrl = (): string => {
  return process.env.REACT_APP_API_BASE_URL || '';
};

/**
 * Make a GET request to the API
 * @param endpoint The API endpoint (without the base URL)
 * @param options Additional fetch options
 * @returns A promise that resolves to the response data
 */
export const apiGet = async (endpoint: string, options?: RequestInit): Promise<any> => {
  const baseUrl = getApiBaseUrl();
  const url = `${baseUrl}${endpoint}`;
  
  try {
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      ...options,
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error(`API GET error for ${endpoint}:`, error);
    throw error;
  }
};

/**
 * Make a POST request to the API
 * @param endpoint The API endpoint (without the base URL)
 * @param data The data to send in the request body
 * @param options Additional fetch options
 * @returns A promise that resolves to the response data
 */
export const apiPost = async (endpoint: string, data: any, options?: RequestInit): Promise<any> => {
  const baseUrl = getApiBaseUrl();
  const url = `${baseUrl}${endpoint}`;
  
  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
      ...options,
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error(`API POST error for ${endpoint}:`, error);
    throw error;
  }
};

// Add more methods as needed (PUT, DELETE, etc.)