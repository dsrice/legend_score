package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	"legend_score/controllers/ci"
	"legend_score/controllers/request"
	"legend_score/entities"
	"legend_score/infra/logger"
	"legend_score/usecases/ui"
)

type authControllerImp struct {
	auth ui.AuthUseCase
}

func NewAuthController(auth ui.AuthUseCase) ci.AuthController {
	return &authControllerImp{
		auth: auth,
	}
}

func (ci *authControllerImp) Login(c echo.Context) error {
	logger.Debug("Login Start")
	var req request.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return ErrorResponse(c, ecode.E0001)
	}

	err = c.Validate(&req)
	if err != nil {
		return ErrorResponse(c, ecode.E0001)
	}

	entity := entities.LoginEntity{
		LoginID:  req.LoginID,
		Password: req.Password,
	}

	err = ci.auth.ValidateLogin(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	token, err := ci.auth.Login(c, &entity)

	logger.Debug("Login End")
	return nil
}
