# Usecase Mocks

This directory contains mock implementations of the usecase interfaces using the [testify/mock](https://github.com/stretchr/testify) package.

## Available Mocks

- `AuthUseCase`: Mock implementation of `ui.AuthUseCase`
- `UserUseCase`: Mock implementation of `ui.UserUseCase`
- `GameUseCase`: Mock implementation of `ui.GameUseCase`

## How to Use

### Basic Usage

```go
import (
    "github.com/stretchr/testify/assert"
    mocklib "github.com/stretchr/testify/mock"
    "legend_score/usecases/mock"
    "testing"
)

func TestSomething(t *testing.T) {
    // Create a new mock instance
    authUseCase := new(mock.AuthUseCase)
    
    // Set up expectations
    authUseCase.On("ValidateLogin", mocklib.Anything, loginEntity).Return(nil)
    
    // Use the mock in your test
    // ...
    
    // Verify all expectations were met
    authUseCase.AssertExpectations(t)
}
```

### Example

See `example_test.go` for complete examples of how to use each mock usecase.

## Adding New Mocks

To add a new mock usecase:

1. Create a new file in this directory with the name of the usecase (e.g., `new_usecase.go`)
2. Define a struct that embeds `mock.Mock`
3. Implement all methods of the usecase interface
4. Add a type assertion to ensure the mock implements the interface

Example:

```go
package mock

import (
    "github.com/stretchr/testify/mock"
    "legend_score/usecases/ui"
)

// NewUseCase is a mock implementation of ui.NewUseCase
type NewUseCase struct {
    mock.Mock
}

// Ensure NewUseCase implements ui.NewUseCase
var _ ui.NewUseCase = (*NewUseCase)(nil)

// SomeMethod mocks the SomeMethod method
func (m *NewUseCase) SomeMethod(args ...interface{}) (returnType, error) {
    args := m.Called(args...)
    
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    
    return args.Get(0).(returnType), args.Error(1)
}
```