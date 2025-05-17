import React from 'react';
import { render, screen } from '@testing-library/react';
import AuthenticatedLayout from './AuthenticatedLayout';
import { useLocation } from 'react-router-dom';

// Mock the dependencies
jest.mock('./Header', () => {
  return function MockHeader({ pageName }) {
    return <div data-testid="header">Header: {pageName}</div>;
  };
});

jest.mock('./Sidebar', () => {
  return function MockSidebar() {
    return <div data-testid="sidebar">Sidebar</div>;
  };
});

// Mock the useLocation hook
jest.mock('react-router-dom', () => ({
  useLocation: jest.fn()
}));

describe('AuthenticatedLayout Component', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('renders layout with all components', () => {
    // Mock the location for home route
    (useLocation as jest.Mock).mockReturnValue({ pathname: '/home' });

    render(
      <AuthenticatedLayout>
        <div data-testid="content">Test Content</div>
      </AuthenticatedLayout>
    );

    // Check if all components are rendered
    expect(screen.getByTestId('sidebar')).toBeInTheDocument();
    expect(screen.getByTestId('header')).toBeInTheDocument();
    expect(screen.getByTestId('content')).toBeInTheDocument();
    
    // Check if the header has the correct page name
    expect(screen.getByTestId('header')).toHaveTextContent('Header: Home');
  });

  test('sets page name to "User List" for /users route', () => {
    // Mock the location for users route
    (useLocation as jest.Mock).mockReturnValue({ pathname: '/users' });

    render(
      <AuthenticatedLayout>
        <div>Test Content</div>
      </AuthenticatedLayout>
    );

    // Check if the header has the correct page name
    expect(screen.getByTestId('header')).toHaveTextContent('Header: User List');
  });

  test('sets default page name for unknown routes', () => {
    // Mock the location for an unknown route
    (useLocation as jest.Mock).mockReturnValue({ pathname: '/unknown' });

    render(
      <AuthenticatedLayout>
        <div>Test Content</div>
      </AuthenticatedLayout>
    );

    // Check if the header has the default page name
    expect(screen.getByTestId('header')).toHaveTextContent('Header: Legend Score');
  });

  test('renders children correctly', () => {
    // Mock the location
    (useLocation as jest.Mock).mockReturnValue({ pathname: '/home' });

    // Render with multiple children
    render(
      <AuthenticatedLayout>
        <div data-testid="child1">Child 1</div>
        <div data-testid="child2">Child 2</div>
      </AuthenticatedLayout>
    );

    // Check if all children are rendered
    expect(screen.getByTestId('child1')).toBeInTheDocument();
    expect(screen.getByTestId('child2')).toBeInTheDocument();
  });
});