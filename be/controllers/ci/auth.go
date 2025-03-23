package ci

import "github.com/labstack/echo/v4"

type AuthController interface {
	Login(c echo.Context) error
}