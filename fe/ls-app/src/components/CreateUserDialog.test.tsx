import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import CreateUserDialog from './CreateUserDialog';
import * as apiClient from '../services/apiClient';
import { useTranslation } from 'react-i18next';

// Mock the apiClient
jest.mock('../services/apiClient', () => ({
  apiPost: jest.fn()
}));

// Mock the useTranslation hook
jest.mock('react-i18next', () => ({
  useTranslation: jest.fn().mockReturnValue({
    t: jest.fn(key => {
      const translations = {
        'createUserDialog.title': 'Create User',
        'createUserDialog.loginId': 'Login ID',
        'createUserDialog.name': 'Name',
        'createUserDialog.password': 'Password',
        'createUserDialog.cancel': 'Cancel',
        'createUserDialog.create': 'Create',
        'createUserDialog.creating': 'Creating...',
        'createUserDialog.error.allFieldsRequired': 'All fields are required',
        'createUserDialog.error.createFailed': 'Failed to create user',
        'createUserDialog.error.generalError': 'An error occurred'
      };
      return translations[key] || key;
    })
  })
}));

describe('CreateUserDialog Component', () => {
  const mockOnClose = jest.fn();
  const mockOnUserCreated = jest.fn();
  
  beforeEach(() => {
    jest.clearAllMocks();
  });
  
  test('renders dialog when isOpen is true', () => {
    render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Check if the dialog title is rendered
    expect(screen.getByText('Create User')).toBeInTheDocument();
    
    // Check if form fields are rendered
    expect(screen.getByLabelText('Login ID')).toBeInTheDocument();
    expect(screen.getByLabelText('Name')).toBeInTheDocument();
    expect(screen.getByLabelText('Password')).toBeInTheDocument();
    
    // Check if buttons are rendered
    expect(screen.getByRole('button', { name: 'Cancel' })).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Create' })).toBeInTheDocument();
  });
  
  test('does not render dialog when isOpen is false', () => {
    render(
      <CreateUserDialog 
        isOpen={false} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Check that the dialog is not rendered
    expect(screen.queryByText('Create User')).not.toBeInTheDocument();
  });
  
  test('handles input changes correctly', () => {
    render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Get form inputs
    const loginIdInput = screen.getByLabelText('Login ID');
    const nameInput = screen.getByLabelText('Name');
    const passwordInput = screen.getByLabelText('Password');
    
    // Change input values
    fireEvent.change(loginIdInput, { target: { value: 'testuser' } });
    fireEvent.change(nameInput, { target: { value: 'Test User' } });
    fireEvent.change(passwordInput, { target: { value: 'password123' } });
    
    // Check if input values are updated
    expect(loginIdInput).toHaveValue('testuser');
    expect(nameInput).toHaveValue('Test User');
    expect(passwordInput).toHaveValue('password123');
  });
  
  test('shows validation error when submitting with empty fields', async () => {
    render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Submit the form without filling in any fields
    fireEvent.click(screen.getByRole('button', { name: 'Create' }));
    
    // Check if validation error is displayed
    await waitFor(() => {
      expect(screen.getByText('All fields are required')).toBeInTheDocument();
    });
    
    // Check that API was not called
    expect(apiClient.apiPost).not.toHaveBeenCalled();
  });
  
  test('submits form successfully', async () => {
    // Mock successful API response
    (apiClient.apiPost as jest.Mock).mockResolvedValue({
      result: true
    });
    
    render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Fill in the form
    fireEvent.change(screen.getByLabelText('Login ID'), { target: { value: 'testuser' } });
    fireEvent.change(screen.getByLabelText('Name'), { target: { value: 'Test User' } });
    fireEvent.change(screen.getByLabelText('Password'), { target: { value: 'password123' } });
    
    // Submit the form
    fireEvent.click(screen.getByRole('button', { name: 'Create' }));
    
    // Wait for the API call to complete
    await waitFor(() => {
      // Check if API was called with correct data
      expect(apiClient.apiPost).toHaveBeenCalledWith('/user', {
        login_id: 'testuser',
        name: 'Test User',
        password: 'password123'
      });
      
      // Check if onClose and onUserCreated were called
      expect(mockOnClose).toHaveBeenCalled();
      expect(mockOnUserCreated).toHaveBeenCalled();
    });
  });
  
  test('handles API error response', async () => {
    // Mock API response with error
    (apiClient.apiPost as jest.Mock).mockResolvedValue({
      result: false,
      message: 'User already exists'
    });
    
    render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Fill in the form
    fireEvent.change(screen.getByLabelText('Login ID'), { target: { value: 'testuser' } });
    fireEvent.change(screen.getByLabelText('Name'), { target: { value: 'Test User' } });
    fireEvent.change(screen.getByLabelText('Password'), { target: { value: 'password123' } });
    
    // Submit the form
    fireEvent.click(screen.getByRole('button', { name: 'Create' }));
    
    // Wait for the API call to complete
    await waitFor(() => {
      // Check if error message is displayed
      expect(screen.getByText('User already exists')).toBeInTheDocument();
      
      // Check that onClose and onUserCreated were not called
      expect(mockOnClose).not.toHaveBeenCalled();
      expect(mockOnUserCreated).not.toHaveBeenCalled();
    });
  });
  
  test('handles API exception', async () => {
    // Mock API call to throw an error
    const mockError = new Error('Network error');
    mockError.response = { data: { message: 'Network error occurred' } };
    (apiClient.apiPost as jest.Mock).mockRejectedValue(mockError);
    
    render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Fill in the form
    fireEvent.change(screen.getByLabelText('Login ID'), { target: { value: 'testuser' } });
    fireEvent.change(screen.getByLabelText('Name'), { target: { value: 'Test User' } });
    fireEvent.change(screen.getByLabelText('Password'), { target: { value: 'password123' } });
    
    // Submit the form
    fireEvent.click(screen.getByRole('button', { name: 'Create' }));
    
    // Wait for the API call to complete
    await waitFor(() => {
      // Check if error message is displayed
      expect(screen.getByText('Network error occurred')).toBeInTheDocument();
      
      // Check that onClose and onUserCreated were not called
      expect(mockOnClose).not.toHaveBeenCalled();
      expect(mockOnUserCreated).not.toHaveBeenCalled();
    });
  });
  
  test('closes dialog when Cancel button is clicked', () => {
    render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Click the Cancel button
    fireEvent.click(screen.getByRole('button', { name: 'Cancel' }));
    
    // Check if onClose was called
    expect(mockOnClose).toHaveBeenCalled();
  });
  
  test('resets form when dialog is opened', () => {
    const { rerender } = render(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Fill in the form
    fireEvent.change(screen.getByLabelText('Login ID'), { target: { value: 'testuser' } });
    fireEvent.change(screen.getByLabelText('Name'), { target: { value: 'Test User' } });
    fireEvent.change(screen.getByLabelText('Password'), { target: { value: 'password123' } });
    
    // Close and reopen the dialog
    rerender(
      <CreateUserDialog 
        isOpen={false} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    rerender(
      <CreateUserDialog 
        isOpen={true} 
        onClose={mockOnClose} 
        onUserCreated={mockOnUserCreated} 
      />
    );
    
    // Check if form fields are reset
    expect(screen.getByLabelText('Login ID')).toHaveValue('');
    expect(screen.getByLabelText('Name')).toHaveValue('');
    expect(screen.getByLabelText('Password')).toHaveValue('');
  });
});