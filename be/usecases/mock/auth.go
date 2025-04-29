package mock

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"legend_score/entities"
	"legend_score/usecases/ui"
)

// AuthUseCase is a mock implementation of ui.AuthUseCase
type AuthUseCase struct {
	mock.Mock
}

// Ensure AuthUseCase implements ui.AuthUseCase
var _ ui.AuthUseCase = (*AuthUseCase)(nil)

// ValidateLogin mocks the ValidateLogin method
func (m *AuthUseCase) ValidateLogin(c echo.Context, entity *entities.LoginEntity) error {
	args := m.Called(c, entity)
	return args.Error(0)
}

// ValidatePassword mocks the ValidatePassword method
func (m *AuthUseCase) ValidatePassword(password string) bool {
	args := m.Called(password)
	return args.Bool(0)
}

// Login mocks the Login method
func (m *AuthUseCase) Login(c echo.Context, e *entities.LoginEntity) (*string, error) {
	args := m.Called(c, e)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*string), args.Error(1)
}