package ui

import (
	"github.com/labstack/echo/v4"
	"legend_score/entities"
)

type GameUseCase interface {
	GetGames(c echo.Context, e *entities.GetGamesEntity) error
	GetGamesByUserID(c echo.Context, e *entities.GetGamesByUserIDEntity) error
	GetGameWithDetails(c echo.Context, e *entities.GetGameWithDetailsEntity) error
}