package repositories

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"legend_score/infra/database/connection"
	"legend_score/infra/database/models"
	"legend_score/infra/logger"
	"legend_score/repositories/ri"
	"time"
)

type userRepository struct {
	con *sql.DB
}

func NewUserRepository(con *connection.Connection) ri.UserRepository {

	return &userRepository{con: con.Conn}
}

func (r *userRepository) Get(c echo.Context, condition []qm.QueryMod) (models.UserSlice, error) {
	logger.Debug("Get start")
	results, err := models.Users(condition...).All(c.Request().Context(), r.con)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug("Get end")
	return results, nil
}

func (r *userRepository) GetLoginID(c echo.Context, loginID string) (*models.User, error) {
	logger.Debug("GetLoginID start")
	mods := []qm.QueryMod{
		models.UserWhere.LoginID.EQ(loginID),
	}

	results, err := models.Users(mods...).All(c.Request().Context(), r.con)

	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("not found")
	}

	logger.Debug("GetLoginID end")
	return results[0], nil
}

func (r *userRepository) Insert(c echo.Context, ut *models.User) error {
	logger.Debug("Insert user_token start")
	ut.CreatedAt = time.Now()
	ut.UpdatedAt = time.Now()
	err := ut.Insert(c.Request().Context(), r.con, boil.Infer())
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Debug("Insert user_token end")
	return nil
}