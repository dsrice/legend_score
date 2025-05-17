package mock

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"legend_score/infra/database/models"
	"legend_score/repositories/ri"
)

// GameRepository is a mock implementation of ri.GameRepository
type GameRepository struct {
	mock.Mock
}

// Ensure GameRepository implements ri.GameRepository
var _ ri.GameRepository = (*GameRepository)(nil)

// GetByUserID mocks the GetByUserID method
func (m *GameRepository) GetByUserID(c echo.Context, userID int) ([]*models.Game, error) {
	args := m.Called(c, userID)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).([]*models.Game), args.Error(1)
}

// GetWithDetails mocks the GetWithDetails method
func (m *GameRepository) GetWithDetails(c echo.Context, gameID int) (*models.Game, error) {
	args := m.Called(c, gameID)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*models.Game), args.Error(1)
}