package usecases

import (
	"encoding/base64"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/scrypt"
	"legend_score/consts/ecode"
	"legend_score/entities"
	"legend_score/entities/db"
	"legend_score/infra/database/models"
	"legend_score/infra/logger"
	"legend_score/repositories/ri"
	"legend_score/usecases/ui"
)

type userUseCase struct {
	user ri.UserRepository
	auth ui.AuthUseCase
}

func NewUserUseCase(user ri.UserRepository, auth ui.AuthUseCase) ui.UserUseCase {
	return &userUseCase{
		user: user,
		auth: auth,
	}
}

func (uc *userUseCase) ValidateCreateUser(c echo.Context, e *entities.CreateUserEntity) error {
	logger.Debug("ValidateCreateUser start")

	condition := []qm.QueryMod{
		models.UserWhere.LoginID.EQ(e.LoginID),
	}

	users, err := uc.user.Get(c, condition)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}

	if len(users) > 0 {
		logger.Error("login_id is used")
		e.Code = ecode.E2001
		return errors.New("login_id is used")
	}

	if !uc.auth.ValidatePassword(e.Password) {
		logger.Error("password does not meet requirements")
		e.Code = ecode.E2002
		return errors.New("password does not meet requirements")
	}

	logger.Debug("ValidateCreateUser end")
	return nil
}

func (uc *userUseCase) CreateUser(c echo.Context, e *entities.CreateUserEntity) error {
	logger.Debug("CreateUser start")
	salt := "legend_score_salt_dev"
	dk, err := scrypt.Key([]byte(e.Password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}
	user := models.User{
		LoginID:        e.LoginID,
		Name:           e.Name,
		Password:       base64.StdEncoding.EncodeToString(dk),
		ChangePassFlag: true,
	}

	err = uc.user.Insert(c, &user)
	if err != nil {
		logger.Error("failed insert user")
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}

	logger.Debug("CreateUser end")
	return nil
}

func (uc *userUseCase) GetUsers(c echo.Context, e *entities.GetUsersEntity) error {
	logger.Debug("GetUsers start")

	// Build query conditions based on filters
	conditions := []qm.QueryMod{}

	if e.UserID != nil {
		conditions = append(conditions, models.UserWhere.ID.EQ(*e.UserID))
	}

	if e.LoginID != nil && *e.LoginID != "" {
		conditions = append(conditions, models.UserWhere.LoginID.EQ(*e.LoginID))
	}

	if e.Name != nil && *e.Name != "" {
		conditions = append(conditions, models.UserWhere.Name.LIKE("%"+*e.Name+"%"))
	}

	// Get users from repository
	users, err := uc.user.Get(c, conditions)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}

	// Convert to entity
	userEntities := make([]db.UserEntity, len(users))
	for i, user := range users {
		var userEntity db.UserEntity
		userEntity.SetEntity(user)
		userEntities[i] = userEntity
	}

	e.Users = userEntities

	logger.Debug("GetUsers end")
	return nil
}

func (uc *userUseCase) GetUser(c echo.Context, e *entities.GetUserEntity) error {
	logger.Debug("GetUser start")

	// Build query condition for user ID
	conditions := []qm.QueryMod{
		models.UserWhere.ID.EQ(e.UserID),
	}

	// Get user from repository
	users, err := uc.user.Get(c, conditions)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}

	// Check if user exists
	if len(users) == 0 {
		logger.Error("user not found")
		e.Code = ecode.E0001 // Invalid request - user not found
		return errors.New("user not found")
	}

	// Convert to entity
	var userEntity db.UserEntity
	userEntity.SetEntity(users[0])
	e.User = userEntity

	logger.Debug("GetUser end")
	return nil
}