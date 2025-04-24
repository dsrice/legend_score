# Frontend Testing with Jest

This project uses Jest and React Testing Library for testing React components.

## Running Tests

To run all tests:

```bash
npm test
```

To run a specific test file:

```bash
npm test -- -t "Login Component"
```

To run tests with coverage report:

```bash
npm test -- --coverage
```

## Test Structure

Tests are organized alongside the components they test:

- `src/App.test.tsx` - Tests for the main App component
- `src/pages/Login.test.tsx` - Tests for the Login component
- `src/utils/PrivateRoute.test.tsx` - Tests for the PrivateRoute component

## Writing Tests

### Basic Component Test

```tsx
import React from 'react';
import { render, screen } from '@testing-library/react';
import MyComponent from './MyComponent';

test('renders component correctly', () => {
  render(<MyComponent />);
  expect(screen.getByText('Expected Text')).toBeInTheDocument();
});
```

### Testing User Interactions

```tsx
import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import MyComponent from './MyComponent';

test('handles button click', () => {
  render(<MyComponent />);
  fireEvent.click(screen.getByRole('button', { name: /click me/i }));
  expect(screen.getByText('Button was clicked')).toBeInTheDocument();
});
```

### Testing Asynchronous Code

```tsx
import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import MyComponent from './MyComponent';

test('loads data asynchronously', async () => {
  render(<MyComponent />);
  fireEvent.click(screen.getByText('Load Data'));
  
  await waitFor(() => {
    expect(screen.getByText('Data loaded')).toBeInTheDocument();
  });
});
```

### Mocking Dependencies

```tsx
import React from 'react';
import { render, screen } from '@testing-library/react';
import MyComponent from './MyComponent';
import * as apiService from '../services/api';

// Mock the API service
jest.mock('../services/api', () => ({
  fetchData: jest.fn()
}));

test('displays fetched data', async () => {
  // Setup the mock return value
  (apiService.fetchData as jest.Mock).mockResolvedValue({ name: 'Test Data' });
  
  render(<MyComponent />);
  
  await waitFor(() => {
    expect(screen.getByText('Test Data')).toBeInTheDocument();
  });
});
```

## Resources

- [Jest Documentation](https://jestjs.io/docs/getting-started)
- [React Testing Library Documentation](https://testing-library.com/docs/react-testing-library/intro/)
- [Testing Library Cheatsheet](https://testing-library.com/docs/react-testing-library/cheatsheet)