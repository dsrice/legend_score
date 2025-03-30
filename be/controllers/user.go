package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	ci "legend_score/controllers/ci"
	"legend_score/controllers/request"
	"legend_score/controllers/response"
	"legend_score/entities"
	"legend_score/infra/logger"
	"legend_score/usecases/ui"
	"net/http"
)

type userController struct {
	uc ui.UserUseCase
}

func NewUserController(uc ui.UserUseCase) ci.UserController {
	return &userController{
		uc: uc,
	}
}

func (uc *userController) CreateUser(c echo.Context) error {
	logger.Debug("Start CreateUser")
	var req request.CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	err = c.Validate(&req)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	entity := entities.CreateUserEntity{}
	entity.SetEntity(&req)

	err = uc.uc.ValidateCreateUser(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, entity.Code)
	}

	err = uc.uc.CreateUser(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, entity.Code)
	}

	res := response.CreateUserResponse{
		Result: true,
	}

	logger.Debug("End CreateUser")
	return c.JSON(http.StatusOK, res)
}
