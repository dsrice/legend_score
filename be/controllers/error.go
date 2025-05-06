package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	"legend_score/controllers/response"
)

func ErrorResponse(c echo.Context, code string) error {
	if code == "" {
		code = ecode.E9000
	}

	res := response.ErrorResponse{
		Code: code,
	}

	return c.JSON(ecode.ErrorMap[code], res)
}