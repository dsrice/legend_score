package ri

import (
	"github.com/labstack/echo/v4"
	"legend_score/infra/database/models"
)

// GameRepository defines the interface for game-related database operations
type GameRepository interface {
	// GetByUserID retrieves all games for a specific user
	GetByUserID(c echo.Context, userID int) ([]*models.Game, error)

	// GetWithDetails retrieves a game with its frames and throws
	GetWithDetails(c echo.Context, gameID int) (*models.Game, error)
}