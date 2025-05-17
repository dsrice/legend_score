import React from 'react';
import { render, screen } from '@testing-library/react';
import Header from './Header';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

// Mock the dependencies
jest.mock('react-router-dom', () => ({
  useNavigate: jest.fn()
}));

jest.mock('react-i18next', () => ({
  useTranslation: jest.fn().mockReturnValue({
    t: jest.fn(key => {
      const translations = {
        'header.home': 'Home Page',
        'header.userManagement': 'User Management Page'
      };
      return translations[key] || key;
    })
  })
}));

// Mock the LanguageSwitcher component
jest.mock('./LanguageSwitcher', () => {
  return function MockLanguageSwitcher() {
    return <div data-testid="language-switcher">Language Switcher</div>;
  };
});

describe('Header Component', () => {
  const mockNavigate = jest.fn();

  beforeEach(() => {
    jest.clearAllMocks();
    (useNavigate as jest.Mock).mockReturnValue(mockNavigate);
  });

  test('renders header with correct page name', () => {
    render(<Header pageName="Test Page" />);

    // Check if the header title is rendered
    expect(screen.getByRole('heading')).toHaveTextContent('Test Page');

    // Check if the language switcher is rendered
    expect(screen.getByTestId('language-switcher')).toBeInTheDocument();
  });

  test('translates "Home" page name correctly', () => {
    render(<Header pageName="Home" />);

    // Check if the translated page name is rendered
    expect(screen.getByRole('heading')).toHaveTextContent('Home Page');
  });

  test('translates "User Management" page name correctly', () => {
    render(<Header pageName="User Management" />);

    // Check if the translated page name is rendered
    expect(screen.getByRole('heading')).toHaveTextContent('User Management Page');
  });

  test('does not translate unknown page names', () => {
    render(<Header pageName="Unknown Page" />);

    // Check if the original page name is rendered
    expect(screen.getByRole('heading')).toHaveTextContent('Unknown Page');
  });

});