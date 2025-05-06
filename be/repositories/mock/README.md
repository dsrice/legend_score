# Repository Mocks

This directory contains mock implementations of the repository interfaces using the [testify/mock](https://github.com/stretchr/testify) package.

## Available Mocks

- `UserRepository`: Mock implementation of `ri.UserRepository`
- `UserTokenRepository`: Mock implementation of `ri.UserTokenRepository`
- `GameRepository`: Mock implementation of `ri.GameRepository`

## How to Use

### Basic Usage

```go
import (
    "github.com/stretchr/testify/assert"
    mocklib "github.com/stretchr/testify/mock"
    "legend_score/repositories/mock"
    "testing"
)

func TestSomething(t *testing.T) {
    // Create a new mock instance
    userRepo := new(mock.UserRepository)
    
    // Set up expectations
    userRepo.On("GetLoginID", mocklib.Anything, "testuser").Return(user, nil)
    
    // Use the mock in your test
    // ...
    
    // Verify all expectations were met
    userRepo.AssertExpectations(t)
}
```

### Example

See `example_test.go` for complete examples of how to use each mock repository.

## Adding New Mocks

To add a new mock repository:

1. Create a new file in this directory with the name of the repository (e.g., `new_repo.go`)
2. Define a struct that embeds `mock.Mock`
3. Implement all methods of the repository interface
4. Add a type assertion to ensure the mock implements the interface

Example:

```go
package mock

import (
    "github.com/stretchr/testify/mock"
    "legend_score/repositories/ri"
)

// NewRepository is a mock implementation of ri.NewRepository
type NewRepository struct {
    mock.Mock
}

// Ensure NewRepository implements ri.NewRepository
var _ ri.NewRepository = (*NewRepository)(nil)

// SomeMethod mocks the SomeMethod method
func (m *NewRepository) SomeMethod(args ...interface{}) (returnType, error) {
    args := m.Called(args...)
    
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    
    return args.Get(0).(returnType), args.Error(1)
}
```