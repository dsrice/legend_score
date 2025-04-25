package ui

import (
	"github.com/labstack/echo/v4"
	"legend_score/entities"
)

// GameUseCase defines the interface for game-related business logic
type GameUseCase interface {
	// GetGamesByUserID retrieves all games for the current user
	GetGamesByUserID(c echo.Context, userID int) (*entities.GamesEntity, error)

	// GetGameDetails retrieves a game with its frames and throws
	GetGameDetails(c echo.Context, gameID int, userID int) (*entities.GameDetailEntity, error)
}