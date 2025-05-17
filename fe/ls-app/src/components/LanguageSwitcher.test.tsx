import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import LanguageSwitcher from './LanguageSwitcher';
import { useTranslation } from 'react-i18next';

// Mock the useTranslation hook
jest.mock('react-i18next', () => ({
  useTranslation: jest.fn()
}));

describe('LanguageSwitcher Component', () => {
  // Mock implementation of i18n
  const mockChangeLanguage = jest.fn();
  
  beforeEach(() => {
    jest.clearAllMocks();
  });
  
  test('renders language buttons correctly', () => {
    // Mock the useTranslation hook to return English as the current language
    (useTranslation as jest.Mock).mockReturnValue({
      i18n: {
        language: 'en',
        changeLanguage: mockChangeLanguage
      }
    });
    
    render(<LanguageSwitcher />);
    
    // Check if both language buttons are rendered
    expect(screen.getByText('EN')).toBeInTheDocument();
    expect(screen.getByText('JP')).toBeInTheDocument();
  });
  
  test('English button has active style when language is English', () => {
    // Mock the useTranslation hook to return English as the current language
    (useTranslation as jest.Mock).mockReturnValue({
      i18n: {
        language: 'en',
        changeLanguage: mockChangeLanguage
      }
    });
    
    render(<LanguageSwitcher />);
    
    // Check if the English button has the active class
    const enButton = screen.getByText('EN');
    expect(enButton.className).toContain('bg-blue-700');
    
    // Check if the Japanese button has the inactive class
    const jpButton = screen.getByText('JP');
    expect(jpButton.className).toContain('bg-blue-500');
  });
  
  test('Japanese button has active style when language is Japanese', () => {
    // Mock the useTranslation hook to return Japanese as the current language
    (useTranslation as jest.Mock).mockReturnValue({
      i18n: {
        language: 'ja',
        changeLanguage: mockChangeLanguage
      }
    });
    
    render(<LanguageSwitcher />);
    
    // Check if the Japanese button has the active class
    const jpButton = screen.getByText('JP');
    expect(jpButton.className).toContain('bg-blue-700');
    
    // Check if the English button has the inactive class
    const enButton = screen.getByText('EN');
    expect(enButton.className).toContain('bg-blue-500');
  });
  
  test('clicking English button calls changeLanguage with "en"', () => {
    // Mock the useTranslation hook
    (useTranslation as jest.Mock).mockReturnValue({
      i18n: {
        language: 'ja',
        changeLanguage: mockChangeLanguage
      }
    });
    
    render(<LanguageSwitcher />);
    
    // Click the English button
    fireEvent.click(screen.getByText('EN'));
    
    // Check if changeLanguage was called with 'en'
    expect(mockChangeLanguage).toHaveBeenCalledWith('en');
  });
  
  test('clicking Japanese button calls changeLanguage with "ja"', () => {
    // Mock the useTranslation hook
    (useTranslation as jest.Mock).mockReturnValue({
      i18n: {
        language: 'en',
        changeLanguage: mockChangeLanguage
      }
    });
    
    render(<LanguageSwitcher />);
    
    // Click the Japanese button
    fireEvent.click(screen.getByText('JP'));
    
    // Check if changeLanguage was called with 'ja'
    expect(mockChangeLanguage).toHaveBeenCalledWith('ja');
  });
});