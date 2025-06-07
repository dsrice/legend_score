package ri

import (
	"github.com/labstack/echo/v4"
	"legend_score/infra/database/models"
)

// GameRepository defines the interface for game-related database operations
type GameRepository interface {
	// GetAll retrieves all games
	GetAll(c echo.Context) ([]*models.Game, error)

	// GetByUserID retrieves all games for a specific user
	GetByUserID(c echo.Context, userID int) ([]*models.Game, error)

	// GetWithDetails retrieves a game with its frames and throws
	GetWithDetails(c echo.Context, gameID int) (*models.Game, error)

	// GetFrameByGameIDAndFrameCount retrieves a frame by game ID and frame count
	GetFrameByGameIDAndFrameCount(c echo.Context, gameID, frameCount int) (*models.Frame, error)

	// CreateFrame creates a new frame
	CreateFrame(c echo.Context, frame *models.Frame) (int, error)

	// RegisterThrow registers a throw
	RegisterThrow(c echo.Context, throw *models.Throw) error
}