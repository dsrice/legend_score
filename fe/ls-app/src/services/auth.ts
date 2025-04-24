interface LoginCredentials {
  login_id: string;
  password: string;
}

interface LoginResponse {
  result: boolean;
  token: string;
  code?: string;
}

/**
 * Authenticates a user with the provided credentials
 * @param credentials The login credentials (login_id and password)
 * @returns A promise that resolves to the login response
 */
export const login = async (credentials: LoginCredentials): Promise<LoginResponse> => {
  try {
    const response = await fetch('/api/v1/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Login error:', error);
    throw error;
  }
};

/**
 * Stores the authentication token in localStorage
 * @param token The JWT token to store
 */
export const storeToken = (token: string): void => {
  localStorage.setItem('auth_token', token);
};

/**
 * Retrieves the authentication token from localStorage
 * @returns The stored JWT token or null if not found
 */
export const getToken = (): string | null => {
  return localStorage.getItem('auth_token');
};

/**
 * Removes the authentication token from localStorage
 */
export const removeToken = (): void => {
  localStorage.removeItem('auth_token');
};

/**
 * Checks if the user is authenticated
 * @returns True if the user has a token stored, false otherwise
 */
export const isAuthenticated = (): boolean => {
  return !!getToken();
};