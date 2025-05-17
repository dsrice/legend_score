import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import Sidebar from './Sidebar';
import { useNavigate, useLocation } from 'react-router-dom';
import { removeToken } from '../services/auth';
import { useTranslation } from 'react-i18next';

// Mock the dependencies
jest.mock('react-router-dom', () => ({
  useNavigate: jest.fn(),
  useLocation: jest.fn()
}));

jest.mock('../services/auth', () => ({
  removeToken: jest.fn()
}));

jest.mock('react-i18next', () => ({
  useTranslation: jest.fn().mockReturnValue({
    t: jest.fn(key => {
      const translations = {
        'appName': 'Legend Score',
        'navigation.home': 'Home',
        'navigation.userManagement': 'User Management',
        'navigation.logout': 'Logout'
      };
      return translations[key] || key;
    })
  })
}));

describe('Sidebar Component', () => {
  const mockNavigate = jest.fn();
  
  beforeEach(() => {
    jest.clearAllMocks();
    (useNavigate as jest.Mock).mockReturnValue(mockNavigate);
    // Default to home route
    (useLocation as jest.Mock).mockReturnValue({ pathname: '/home' });
  });
  
  test('renders sidebar with app name and navigation items', () => {
    render(<Sidebar />);
    
    // Check if app name is rendered
    expect(screen.getByText('Legend Score')).toBeInTheDocument();
    
    // Check if navigation items are rendered
    expect(screen.getByText('Home')).toBeInTheDocument();
    expect(screen.getByText('User Management')).toBeInTheDocument();
    expect(screen.getByText('Logout')).toBeInTheDocument();
  });
  
  test('home button has active style when on home route', () => {
    // Mock location to be on home route
    (useLocation as jest.Mock).mockReturnValue({ pathname: '/home' });
    
    render(<Sidebar />);
    
    // Get the home button (the button containing the text 'Home')
    const homeButton = screen.getByText('Home').closest('button');
    
    // Check if it has the active class
    expect(homeButton?.className).toContain('bg-blue-600');
    
    // Get the users button
    const usersButton = screen.getByText('User Management').closest('button');
    
    // Check if it doesn't have the active class
    expect(usersButton?.className).not.toContain('bg-blue-600');
    expect(usersButton?.className).toContain('hover:bg-gray-700');
  });
  
  test('users button has active style when on users route', () => {
    // Mock location to be on users route
    (useLocation as jest.Mock).mockReturnValue({ pathname: '/users' });
    
    render(<Sidebar />);
    
    // Get the users button
    const usersButton = screen.getByText('User Management').closest('button');
    
    // Check if it has the active class
    expect(usersButton?.className).toContain('bg-blue-600');
    
    // Get the home button
    const homeButton = screen.getByText('Home').closest('button');
    
    // Check if it doesn't have the active class
    expect(homeButton?.className).not.toContain('bg-blue-600');
    expect(homeButton?.className).toContain('hover:bg-gray-700');
  });
  
  test('clicking home button navigates to home route', () => {
    render(<Sidebar />);
    
    // Click the home button
    fireEvent.click(screen.getByText('Home'));
    
    // Check if navigate was called with the correct path
    expect(mockNavigate).toHaveBeenCalledWith('/home');
  });
  
  test('clicking users button navigates to users route', () => {
    render(<Sidebar />);
    
    // Click the users button
    fireEvent.click(screen.getByText('User Management'));
    
    // Check if navigate was called with the correct path
    expect(mockNavigate).toHaveBeenCalledWith('/users');
  });
  
  test('clicking logout button calls removeToken and navigates to login page', () => {
    render(<Sidebar />);
    
    // Click the logout button
    fireEvent.click(screen.getByText('Logout'));
    
    // Check if removeToken was called
    expect(removeToken).toHaveBeenCalled();
    
    // Check if navigate was called with the correct path
    expect(mockNavigate).toHaveBeenCalledWith('/');
  });
  
  test('translations are applied correctly', () => {
    // Create a custom mock for useTranslation
    (useTranslation as jest.Mock).mockReturnValue({
      t: jest.fn(key => {
        const customTranslations = {
          'appName': 'Custom App Name',
          'navigation.home': 'Custom Home',
          'navigation.userManagement': 'Custom User Management',
          'navigation.logout': 'Custom Logout'
        };
        return customTranslations[key] || key;
      })
    });
    
    render(<Sidebar />);
    
    // Check if the custom translations are applied
    expect(screen.getByText('Custom App Name')).toBeInTheDocument();
    expect(screen.getByText('Custom Home')).toBeInTheDocument();
    expect(screen.getByText('Custom User Management')).toBeInTheDocument();
    expect(screen.getByText('Custom Logout')).toBeInTheDocument();
  });
});