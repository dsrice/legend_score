import React from 'react';
import { render, screen } from '@testing-library/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import PrivateRoute from './PrivateRoute';
import * as authService from '../services/auth';

// Mock the auth service
jest.mock('../services/auth', () => ({
  isAuthenticated: jest.fn()
}));

describe('PrivateRoute Component', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('renders children when user is authenticated', () => {
    // Mock isAuthenticated to return true
    (authService.isAuthenticated as jest.Mock).mockReturnValue(true);
    
    render(
      <MemoryRouter initialEntries={['/protected']}>
        <Routes>
          <Route 
            path="/protected" 
            element={
              <PrivateRoute>
                <div data-testid="protected-content">Protected Content</div>
              </PrivateRoute>
            } 
          />
        </Routes>
      </MemoryRouter>
    );
    
    // Check if the protected content is rendered
    expect(screen.getByTestId('protected-content')).toBeInTheDocument();
    expect(screen.getByText('Protected Content')).toBeInTheDocument();
  });

  test('redirects to login page when user is not authenticated', () => {
    // Mock isAuthenticated to return false
    (authService.isAuthenticated as jest.Mock).mockReturnValue(false);
    
    render(
      <MemoryRouter initialEntries={['/protected']}>
        <Routes>
          <Route path="/" element={<div data-testid="login-page">Login Page</div>} />
          <Route 
            path="/protected" 
            element={
              <PrivateRoute>
                <div data-testid="protected-content">Protected Content</div>
              </PrivateRoute>
            } 
          />
        </Routes>
      </MemoryRouter>
    );
    
    // Check if redirected to login page
    expect(screen.getByTestId('login-page')).toBeInTheDocument();
    expect(screen.getByText('Login Page')).toBeInTheDocument();
    
    // Check that protected content is not rendered
    expect(screen.queryByTestId('protected-content')).not.toBeInTheDocument();
  });
});