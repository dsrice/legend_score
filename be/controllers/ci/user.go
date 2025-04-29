package ci

import "github.com/labstack/echo/v4"

type UserController interface {
	CreateUser(c echo.Context) error
	GetUsers(c echo.Context) error  // New method
}