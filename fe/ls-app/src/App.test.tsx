import React from 'react';
import { render, screen, act } from '@testing-library/react';
import App from './App';

// Mock the env module
jest.mock('./services/env', () => ({
  getApiBaseUrl: jest.fn().mockReturnValue('http://localhost:8080'),
  setApiBaseUrl: jest.fn()
}));

test('renders the App component with router', async () => {
  await act(async () => {
    render(<App />);
  });
  // Check if the Login route is rendered by default
  const loginElement = screen.getByRole('heading', { name: /sign in to your account/i });
  expect(loginElement).toBeInTheDocument();
});