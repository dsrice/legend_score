package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/controllers/ci"
)

type authControllerImp struct{}

func NewAuth() ci.Auth {
	return &authControllerImp{}
}

func (ci *authControllerImp) Login(c echo.Context) error {

	return nil
}