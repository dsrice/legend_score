import React from 'react';
import { render, screen } from '@testing-library/react';
import Home from './Home';

describe('Home Component', () => {
  test('renders the welcome message', () => {
    render(<Home />);
    
    // Check if the welcome title is rendered
    expect(screen.getByText('Welcome to Legend Score')).toBeInTheDocument();
    
    // Check if the logged in message is rendered
    expect(screen.getByText('You are successfully logged in!')).toBeInTheDocument();
  });

  test('renders with the correct CSS classes', () => {
    const { container } = render(<Home />);
    
    // Check if the main container has the correct classes
    const mainDiv = container.firstChild;
    expect(mainDiv).toHaveClass('py-12');
    expect(mainDiv).toHaveClass('px-4');
    expect(mainDiv).toHaveClass('sm:px-6');
    expect(mainDiv).toHaveClass('lg:px-8');
    
    // Check if the inner container has the correct classes
    const innerDiv = container.firstChild?.firstChild;
    expect(innerDiv).toHaveClass('max-w-md');
    expect(innerDiv).toHaveClass('mx-auto');
    expect(innerDiv).toHaveClass('space-y-8');
  });

  test('renders the heading with correct styling', () => {
    render(<Home />);
    
    const heading = screen.getByText('Welcome to Legend Score');
    expect(heading).toHaveClass('text-center');
    expect(heading).toHaveClass('text-3xl');
    expect(heading).toHaveClass('font-extrabold');
    expect(heading).toHaveClass('text-gray-900');
  });

  test('renders the paragraph with correct styling', () => {
    render(<Home />);
    
    const paragraph = screen.getByText('You are successfully logged in!');
    expect(paragraph).toHaveClass('mt-2');
    expect(paragraph).toHaveClass('text-center');
    expect(paragraph).toHaveClass('text-sm');
    expect(paragraph).toHaveClass('text-gray-600');
  });
});