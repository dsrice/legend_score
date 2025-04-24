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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags user
// @Accept json
// @Produce json
// @Param user body request.CreateUserRequest true "User information"
// @Success 200 {object} response.CreateUserResponse
// @Failure 400 {object} response.CreateUserResponse
// @Router /user [post]
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
