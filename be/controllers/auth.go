package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	"legend_score/controllers/ci"
	"legend_score/controllers/request"
	"legend_score/infra/logger"
)

type authControllerImp struct{}

func NewAuth() ci.AuthController {
	return &authControllerImp{}
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

	logger.Debug("Login End")
	return nil
}
