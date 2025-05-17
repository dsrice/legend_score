import React from 'react';
import { render, screen, fireEvent, waitFor, act } from '@testing-library/react';
import CreateUserDialog from './CreateUserDialog';
import * as apiClient from '../services/apiClient';
import { JSX } from 'react/jsx-runtime';

// Mock the apiClient
jest.mock('../services/apiClient', () => ({
  apiPost: jest.fn()
}));

describe('CreateUserDialog Component', () => {
  const mockProps = {
    isOpen: true,
    onClose: jest.fn(),
    onUserCreated: jest.fn()
  };

  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('renders dialog when isOpen is true', async () => {
    await act(async () => {
      render(<CreateUserDialog {...mockProps} />);
    });

    // Check if the dialog title is rendered
    expect(screen.getByText(/create new user/i)).toBeInTheDocument();

    // Check if form inputs are rendered
    expect(screen.getByLabelText(/login id/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/name/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();

    // Check if buttons are rendered
    expect(screen.getByRole('button', {name: /cancel/i})).toBeInTheDocument();
    expect(screen.getByRole('button', {name: /create user/i})).toBeInTheDocument();
  });

  test('does not render dialog when isOpen is false', async () => {
    await act(async () => {
      render(<CreateUserDialog {...mockProps} isOpen={false}/>);
    });

    // Check if the dialog is not rendered
    expect(screen.queryByText(/create new user/i)).not.toBeInTheDocument();
  });

  test('calls onClose when Cancel button is clicked', async () => {
    await act(async () => {
      render(<CreateUserDialog {...mockProps} />);
    });

    // Click the Cancel button
    await act(async () => {
      fireEvent.click(screen.getByRole('button', {name: /cancel/i}));
    });

    // Check if onClose was called
    expect(mockProps.onClose).toHaveBeenCalled();
  });

  test('handles form submission correctly', async () => {
    // Mock successful API response
    (apiClient.apiPost as jest.Mock).mockResolvedValue({
      result: true
    });

    await act(async () => {
      render(<CreateUserDialog {...mockProps} />);
    });

    // Fill in the form
    await act(async () => {
      fireEvent.change(screen.getByLabelText(/login id/i), {
        target: {value: 'testuser'}
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/name/i), {
        target: {value: 'Test User'}
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/password/i), {
        target: {value: 'password123'}
      });
    });

    // Submit the form
    await act(async () => {
      fireEvent.click(screen.getByRole('button', {name: /create user/i}));
    });

    // Wait for the API call to complete
    await waitFor(() => {
      // Check if API was called with correct data
      expect(apiClient.apiPost).toHaveBeenCalledWith('/users', {
        login_id: 'testuser',
        name: 'Test User',
        password: 'password123'
      });

      // Check if onClose and onUserCreated were called
      expect(mockProps.onClose).toHaveBeenCalled();
      expect(mockProps.onUserCreated).toHaveBeenCalled();
    });
  });

  test('displays error message on API failure', async () => {
    // Mock failed API response
    (apiClient.apiPost as jest.Mock).mockResolvedValue({
      result: false,
      message: 'User already exists'
    });

    await act(async () => {
      render(<CreateUserDialog {...mockProps} />);
    });

    // Fill in the form
    await act(async () => {
      fireEvent.change(screen.getByLabelText(/login id/i), {
        target: {value: 'existinguser'}
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/name/i), {
        target: {value: 'Existing User'}
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/password/i), {
        target: {value: 'password123'}
      });
    });

    // Submit the form
    await act(async () => {
      fireEvent.click(screen.getByRole('button', {name: /create user/i}));
    });

    // Wait for the error message to appear
    await waitFor(() => {
      expect(screen.getByText(/user already exists/i)).toBeInTheDocument();
    });

    // Check that onClose and onUserCreated were not called
    expect(mockProps.onClose).not.toHaveBeenCalled();
    expect(mockProps.onUserCreated).not.toHaveBeenCalled();
  });

  test('handles API exceptions', async () => {
    // Mock API exception
    (apiClient.apiPost as jest.Mock).mockRejectedValue({
      response: {
        data: {
          message: 'Network error'
        }
      }
    });

    await act(async () => {
      render(<CreateUserDialog {...mockProps} />);
    });

    // Fill in the form
    await act(async () => {
      fireEvent.change(screen.getByLabelText(/login id/i), {
        target: {value: 'testuser'}
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/name/i), {
        target: {value: 'Test User'}
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/password/i), {
        target: {value: 'password123'}
      });
    });

    // Submit the form
    await act(async () => {
      fireEvent.click(screen.getByRole('button', {name: /create user/i}));
    });

    // Wait for the error message to appear
    await waitFor(() => {
      expect(screen.getByText(/network error/i)).toBeInTheDocument();
    });
  });

  test.skip('resets form when dialog is opened', async () => {
    let rerenderFn: ((ui: React.ReactNode) => void) | ((arg0: JSX.Element) => void);

    await act(async () => {
      const { rerender } = render(<CreateUserDialog {...mockProps} isOpen={false} />);
      rerenderFn = rerender;
    });

    // Re-render with isOpen=true
    await act(async () => {
      rerenderFn(<CreateUserDialog {...mockProps} isOpen={true} />);
    });

    // Check if form inputs are empty
    expect(screen.getByLabelText(/login id/i)).toHaveValue('');
    expect(screen.getByLabelText(/name/i)).toHaveValue('');
    expect(screen.getByLabelText(/password/i)).toHaveValue('');
  });
});