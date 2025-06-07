package ci

import "github.com/labstack/echo/v4"

type GameController interface {
    GetGames(c echo.Context) error
    GetGamesByUserID(c echo.Context) error
    GetGameWithDetails(c echo.Context) error
    RegisterThrow(c echo.Context) error
}