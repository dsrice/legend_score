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

// GetGamesByUserID mocks the GetGamesByUserID method
func (m *GameUseCase) GetGamesByUserID(c echo.Context, userID int) (*entities.GamesEntity, error) {
	args := m.Called(c, userID)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*entities.GamesEntity), args.Error(1)
}

// GetGameDetails mocks the GetGameDetails method
func (m *GameUseCase) GetGameDetails(c echo.Context, gameID int, userID int) (*entities.GameDetailEntity, error) {
	args := m.Called(c, gameID, userID)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*entities.GameDetailEntity), args.Error(1)
}