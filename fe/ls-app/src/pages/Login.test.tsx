import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Login from './Login';
import * as authService from '../services/auth';

// Mock the auth service
jest.mock('../services/auth', () => ({
  login: jest.fn(),
  storeToken: jest.fn(),
  isAuthenticated: jest.fn()
}));

// Mock the useNavigate hook
const mockNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate
}));

describe('Login Component', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('renders login form correctly', () => {
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );
    
    // Check if the heading is rendered
    expect(screen.getByRole('heading', { name: /sign in to your account/i })).toBeInTheDocument();
    
    // Check if form inputs are rendered
    expect(screen.getByPlaceholderText(/login id/i)).toBeInTheDocument();
    expect(screen.getByPlaceholderText(/password/i)).toBeInTheDocument();
    
    // Check if the submit button is rendered
    expect(screen.getByRole('button', { name: /sign in/i })).toBeInTheDocument();
  });

  test('redirects to home if already authenticated', () => {
    // Mock isAuthenticated to return true
    (authService.isAuthenticated as jest.Mock).mockReturnValue(true);
    
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );
    
    // Check if navigate was called with the correct path
    expect(mockNavigate).toHaveBeenCalledWith('/home');
  });

  test('handles form submission correctly', async () => {
    // Mock successful login
    (authService.login as jest.Mock).mockResolvedValue({
      result: true,
      token: 'fake-token'
    });
    
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );
    
    // Fill in the form
    fireEvent.change(screen.getByPlaceholderText(/login id/i), {
      target: { value: 'testuser' }
    });
    
    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: 'password123' }
    });
    
    // Submit the form
    fireEvent.click(screen.getByRole('button', { name: /sign in/i }));
    
    // Wait for the login process to complete
    await waitFor(() => {
      // Check if login was called with correct data
      expect(authService.login).toHaveBeenCalledWith({
        login_id: 'testuser',
        password: 'password123'
      });
      
      // Check if token was stored
      expect(authService.storeToken).toHaveBeenCalledWith('fake-token');
      
      // Check if navigation occurred
      expect(mockNavigate).toHaveBeenCalledWith('/home');
    });
  });

  test('displays error message on login failure', async () => {
    // Mock failed login
    (authService.login as jest.Mock).mockResolvedValue({
      result: false,
      code: 'Invalid credentials'
    });
    
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );
    
    // Fill in the form
    fireEvent.change(screen.getByPlaceholderText(/login id/i), {
      target: { value: 'testuser' }
    });
    
    fireEvent.change(screen.getByPlaceholderText(/password/i), {
      target: { value: 'wrongpassword' }
    });
    
    // Submit the form
    fireEvent.click(screen.getByRole('button', { name: /sign in/i }));
    
    // Wait for the error message to appear
    await waitFor(() => {
      expect(screen.getByText(/invalid credentials/i)).toBeInTheDocument();
    });
  });
});