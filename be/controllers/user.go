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

// GetUsers godoc
// @Summary Get a list of users
// @Description Get a list of users with optional filtering by user ID, login ID, and name
// @Tags user
// @Accept json
// @Produce json
// @Param user_id query int false "Filter by user ID"
// @Param login_id query string false "Filter by login ID"
// @Param name query string false "Filter by user name"
// @Success 200 {object} response.GetUsersResponse
// @Failure 400 {object} response.GetUsersResponse
// @Router /users [get]
func (uc *userController) GetUsers(c echo.Context) error {
	logger.Debug("Start GetUsers")

	// Parse query parameters
	var req request.GetUsersRequest
	if err := c.Bind(&req); err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	// Create entity and set filters
	entity := entities.GetUsersEntity{}
	entity.SetFilters(req.UserID, req.LoginID, req.Name)

	// Call use case
	err := uc.uc.GetUsers(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, entity.Code)
	}

	// Create response
	users := make([]response.UserResponse, len(entity.Users))
	for i, user := range entity.Users {
		users[i] = response.UserResponse{
			ID:      user.ID,
			LoginID: user.LoginID,
			Name:    user.Name,
		}
	}

	res := response.GetUsersResponse{
		Result: true,
		Users:  users,
	}

	logger.Debug("End GetUsers")
	return c.JSON(http.StatusOK, res)
}