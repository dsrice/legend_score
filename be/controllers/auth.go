package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	"legend_score/controllers/ci"
	"legend_score/controllers/request"
	"legend_score/controllers/response"
	"legend_score/entities"
	"legend_score/infra/logger"
	"legend_score/usecases/ui"
	"net/http"
)

type authControllerImp struct {
	auth ui.AuthUseCase
}

func NewAuthController(auth ui.AuthUseCase) ci.AuthController {
	return &authControllerImp{
		auth: auth,
	}
}

// Login godoc
// @Summary Login to the application
// @Description Authenticate user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login_request body request.LoginRequest true "Login credentials"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.LoginResponse
// @Router /login [post]
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

	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	res := response.LoginResponse{
		Token:  *token,
		Result: true,
	}
	logger.Debug("Login End")
	return c.JSON(http.StatusOK, res)
}
