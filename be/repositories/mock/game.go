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

// GetAll mocks the GetAll method
func (m *GameRepository) GetAll(c echo.Context) ([]*models.Game, error) {
	args := m.Called(c)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*models.Game), args.Error(1)
}

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

// GetFrameByGameIDAndFrameCount mocks the GetFrameByGameIDAndFrameCount method
func (m *GameRepository) GetFrameByGameIDAndFrameCount(c echo.Context, gameID, frameCount int) (*models.Frame, error) {
	args := m.Called(c, gameID, frameCount)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Frame), args.Error(1)
}

// CreateFrame mocks the CreateFrame method
func (m *GameRepository) CreateFrame(c echo.Context, frame *models.Frame) (int, error) {
	args := m.Called(c, frame)
	return args.Int(0), args.Error(1)
}

// RegisterThrow mocks the RegisterThrow method
func (m *GameRepository) RegisterThrow(c echo.Context, throw *models.Throw) error {
	args := m.Called(c, throw)
	return args.Error(0)
}