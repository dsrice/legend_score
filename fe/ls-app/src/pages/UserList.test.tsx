import React from 'react';
import { render, screen, fireEvent, waitFor, act } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import UserList from './UserList';
import * as apiClient from '../services/apiClient';

// Mock the env module
jest.mock('../services/env', () => ({
  getApiBaseUrl: jest.fn().mockReturnValue('http://localhost:8080'),
  setApiBaseUrl: jest.fn()
}));

// Mock the apiClient
jest.mock('../services/apiClient', () => ({
  apiGet: jest.fn()
}));

// Mock the CreateUserDialog component
jest.mock('../components/CreateUserDialog', () => {
  return function MockCreateUserDialog({ isOpen, onClose, onUserCreated }) {
    return isOpen ? (
      <div data-testid="create-user-dialog">
        <button onClick={onClose}>Close</button>
        <button onClick={onUserCreated}>register</button>
      </div>
    ) : null;
  };
});

// Mock the useNavigate hook
const mockNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate
}));

describe('UserList Component', () => {
  const mockUsers = [
    { id: 1, login_id: 'user1', name: 'User One' },
    { id: 2, login_id: 'user2', name: 'User Two' }
  ];

  beforeEach(() => {
    jest.clearAllMocks();
    // Use fake timers to control asynchronous operations
    jest.useFakeTimers();
    // Mock successful API response by default
    (apiClient.apiGet as jest.Mock).mockResolvedValue({
      result: true,
      users: mockUsers
    });
  });

  afterEach(() => {
    // Restore real timers
    jest.useRealTimers();
  });

  test('renders user list page correctly', async () => {
    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Check if the heading is rendered
    expect(screen.getByRole('heading', { name: /user list/i })).toBeInTheDocument();

    // Check if the filter form is rendered
    expect(screen.getByRole('heading', { name: /filter users/i })).toBeInTheDocument();

    // Check if the buttons are rendered
    expect(screen.getByRole('button', { name: /create user/i })).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /back to home/i })).toBeInTheDocument();

    // Wait for the users to be loaded
    await waitFor(() => {
      expect(apiClient.apiGet).toHaveBeenCalledWith('/user', {});
    });
  });

  test('displays users in the table', async () => {
    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Wait for the users to be loaded
    await waitFor(() => {
      // Check if the table headers are rendered
      const tableHeaders = screen.getAllByRole('columnheader');
      expect(tableHeaders.length).toBe(3);
      expect(tableHeaders[0]).toHaveTextContent(/id/i);
      expect(tableHeaders[1]).toHaveTextContent(/login id/i);
      expect(tableHeaders[2]).toHaveTextContent(/name/i);

      // Check if the user data is rendered
      expect(screen.getByText('user1')).toBeInTheDocument();
      expect(screen.getByText('User One')).toBeInTheDocument();
      expect(screen.getByText('user2')).toBeInTheDocument();
      expect(screen.getByText('User Two')).toBeInTheDocument();
    });
  });

  test('displays loading state while fetching users', async () => {
    // Delay the API response
    (apiClient.apiGet as jest.Mock).mockImplementation(() => {
      return new Promise(resolve => {
        setTimeout(() => {
          resolve({
            result: true,
            users: mockUsers
          });
        }, 100);
      });
    });

    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Check if loading message is displayed
    expect(screen.getByText(/loading users/i)).toBeInTheDocument();

    // Wait for the users to be loaded
    await waitFor(() => {
      expect(screen.queryByText(/loading users/i)).not.toBeInTheDocument();
    });
  });

  test('displays error message when API request fails', async () => {
    // Mock failed API response
    (apiClient.apiGet as jest.Mock).mockResolvedValue({
      result: false
    });

    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Wait for the error message to be displayed
    await waitFor(() => {
      expect(screen.getByText(/failed to fetch users/i)).toBeInTheDocument();
    });
  });

  test('navigates back to home when Back to Home button is clicked', async () => {
    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Click the Back to Home button
    fireEvent.click(screen.getByRole('button', { name: /back to home/i }));

    // Check if navigate was called with the correct path
    expect(mockNavigate).toHaveBeenCalledWith('/home');
  });

  test('opens create user dialog when Create User button is clicked', async () => {
    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Click the Create User button
    await act(async () => {
      fireEvent.click(screen.getByRole('button', { name: /create user/i }));
    });

    // Check if the dialog is opened
    await waitFor(() => {
      expect(screen.getByTestId('create-user-dialog')).toBeInTheDocument();
    });
  });

  test('applies filters when filter form is submitted', async () => {
    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Fill in the filter form
    await act(async () => {
      fireEvent.change(screen.getByLabelText(/user id/i), {
        target: { value: '1' }
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/login id/i), {
        target: { value: 'test' }
      });
    });

    await act(async () => {
      fireEvent.change(screen.getByLabelText(/name/i), {
        target: { value: 'Test User' }
      });
    });

    // Submit the form
    await act(async () => {
      fireEvent.click(screen.getByRole('button', { name: /apply filters/i }));
    });

    // Check if the API was called with the correct filters
    await waitFor(() => {
      expect(apiClient.apiGet).toHaveBeenCalledWith('/user', {
        user_id: '1',
        login_id: 'test',
        name: 'Test User'
      });
    });
  });

  test('resets filters when Reset button is clicked', async () => {
    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Fill in the filter form
    await act(async () => {
      fireEvent.change(screen.getByLabelText(/user id/i), {
        target: { value: '1' }
      });
    });

    // Click the Reset button
    await act(async () => {
      fireEvent.click(screen.getByRole('button', { name: /reset/i }));
    });

    // Check if the input fields are cleared
    expect(screen.getByLabelText(/user id/i)).toHaveValue('');

    // Check if the API was called without filters
    await waitFor(() => {
      expect(apiClient.apiGet).toHaveBeenCalledWith('/user', {});
    });
  });

  test('refreshes user list when a user is created', async () => {
    await act(async () => {
      render(
        <BrowserRouter>
          <UserList />
        </BrowserRouter>
      );
    });

    // Clear previous API calls
    (apiClient.apiGet as jest.Mock).mockClear();

    // Click the Create User button to open the dialog
    await act(async () => {
      fireEvent.click(screen.getByRole('button', { name: /create user/i }));
    });

    // Simulate user creation
    await act(async () => {
      fireEvent.click(screen.getByText(/register/i));
    });

    // Advance timers to trigger any pending timeouts
    await act(async () => {
      jest.advanceTimersByTime(100);
    });

    // Run all pending promises
    await act(async () => {
      await Promise.resolve();
    });

    // Check if the API was called to refresh the user list
    expect(apiClient.apiGet).toHaveBeenCalledWith('/user', {});
  });
});