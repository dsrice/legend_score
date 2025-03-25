package repositories

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"legend_score/infra/database/connection"
	"legend_score/infra/database/models"
	"legend_score/infra/logger"
	"legend_score/repositories/ri"
	"time"
)

type userTokenRepository struct {
	con *sql.DB
}

func NewUserTokenRepository(con *connection.Connection) ri.UserTokenRepository {
	return &userTokenRepository{con: con.Conn}
}

func (r *userTokenRepository) Insert(c echo.Context, ut *models.UserToken) error {
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