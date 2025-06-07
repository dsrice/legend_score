package mock

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"legend_score/entities"
	"legend_score/usecases/ui"
)

// GameUseCase is a mock implementation of ui.GameUseCase
type GameUseCase struct {
	mock.Mock
}

// Ensure GameUseCase implements ui.GameUseCase
var _ ui.GameUseCase = (*GameUseCase)(nil)

// GetGames mocks the GetGames method
func (m *GameUseCase) GetGames(c echo.Context, e *entities.GetGamesEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}

// GetGamesByUserID mocks the GetGamesByUserID method
func (m *GameUseCase) GetGamesByUserID(c echo.Context, e *entities.GetGamesByUserIDEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}

// GetGameWithDetails mocks the GetGameWithDetails method
func (m *GameUseCase) GetGameWithDetails(c echo.Context, e *entities.GetGameWithDetailsEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}

// RegisterThrow mocks the RegisterThrow method
func (m *GameUseCase) RegisterThrow(c echo.Context, e *entities.RegisterThrowEntity) error {
	args := m.Called(c, e)
	return args.Error(0)
}