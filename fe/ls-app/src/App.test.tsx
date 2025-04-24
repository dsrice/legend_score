import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders the App component with router', () => {
  render(<App />);
  // Check if the Login route is rendered by default
  const loginElement = screen.getByRole('heading', { name: /sign in to your account/i });
  expect(loginElement).toBeInTheDocument();
});