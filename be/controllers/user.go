package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	ci "legend_score/controllers/ci"
	"legend_score/controllers/request"
	"legend_score/entities"
	"legend_score/infra/logger"
)

type userController struct{}

func NewUser() ci.UserController {
	return &userController{}
}

func (uc *userController) CreateUser(c echo.Context) error {
	logger.Debug("Start CreateUser")
	var req request.CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		return ErrorResponse(c, ecode.E0001)
	}

	err = c.Validate(&req)
	if err != nil {
		return ErrorResponse(c, ecode.E0001)
	}

	entity := entities.CreateUserEntity{}
	entity.SetEntity(req)

	logger.Debug("End CreateUser")
	return nil
}