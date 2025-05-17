package mock

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"legend_score/entities"
	"legend_score/usecases/ui"
)

// UserUseCase is a mock implementation of ui.UserUseCase
type UserUseCase struct {
	mock.Mock
}

// Ensure UserUseCase implements ui.UserUseCase
var _ ui.UserUseCase = (*UserUseCase)(nil)

// ValidateCreateUser mocks the ValidateCreateUser method
func (m *UserUseCase) ValidateCreateUser(c echo.Context, e *entities.CreateUserEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}

// CreateUser mocks the CreateUser method
func (m *UserUseCase) CreateUser(c echo.Context, e *entities.CreateUserEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}

// GetUsers mocks the GetUsers method
func (m *UserUseCase) GetUsers(c echo.Context, e *entities.GetUsersEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}

// GetUser mocks the GetUser method
func (m *UserUseCase) GetUser(c echo.Context, e *entities.GetUserEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}