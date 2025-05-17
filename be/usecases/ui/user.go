package ui

import (
	"github.com/labstack/echo/v4"
	"legend_score/entities"
)

type UserUseCase interface {
	ValidateCreateUser(c echo.Context, e *entities.CreateUserEntity) error
	CreateUser(c echo.Context, e *entities.CreateUserEntity) error
	GetUsers(c echo.Context, e *entities.GetUsersEntity) error
	GetUser(c echo.Context, e *entities.GetUserEntity) error
}