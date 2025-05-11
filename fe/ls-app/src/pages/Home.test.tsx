import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Home from './Home';
import * as authService from '../services/auth';

// Mock the auth service
jest.mock('../services/auth', () => ({
  removeToken: jest.fn()
}));

// Mock the useNavigate hook
const mockNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate
}));

describe('Home Component', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('renders welcome message correctly', () => {
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    );
    
    // Check if the heading is rendered
    expect(screen.getByRole('heading', { name: /welcome to legend score/i })).toBeInTheDocument();
    
    // Check if the welcome text is rendered
    expect(screen.getByText(/you are successfully logged in/i)).toBeInTheDocument();
  });

  test('renders navigation buttons correctly', () => {
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    );
    
    // Check if the User List button is rendered
    expect(screen.getByRole('button', { name: /user list/i })).toBeInTheDocument();
    
    // Check if the Logout button is rendered
    expect(screen.getByRole('button', { name: /logout/i })).toBeInTheDocument();
  });

  test('navigates to user list page when User List button is clicked', () => {
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    );
    
    // Click the User List button
    fireEvent.click(screen.getByRole('button', { name: /user list/i }));
    
    // Check if navigate was called with the correct path
    expect(mockNavigate).toHaveBeenCalledWith('/users');
  });

  test('logs out and navigates to login page when Logout button is clicked', () => {
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    );
    
    // Click the Logout button
    fireEvent.click(screen.getByRole('button', { name: /logout/i }));
    
    // Check if removeToken was called
    expect(authService.removeToken).toHaveBeenCalled();
    
    // Check if navigate was called with the correct path
    expect(mockNavigate).toHaveBeenCalledWith('/');
  });
});