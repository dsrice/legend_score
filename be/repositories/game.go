package repositories

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"legend_score/infra/database/connection"
	"legend_score/infra/database/models"
	"legend_score/infra/logger"
	"legend_score/repositories/ri"
)

type gameRepository struct {
	con *sql.DB
}

// NewGameRepository creates a new instance of GameRepository
func NewGameRepository(con *connection.Connection) ri.GameRepository {
	return &gameRepository{con: con.Conn}
}

// GetByUserID retrieves all games for a specific user
func (r *gameRepository) GetByUserID(c echo.Context, userID int) ([]*models.Game, error) {
	logger.Debug("GetByUserID start")
	games, err := models.Games(
		qm.Where("user_id = ?", userID),
		qm.Where("deleted_flg = ?", false),
		qm.OrderBy("game_date DESC"),
	).All(c.Request().Context(), r.con)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug("GetByUserID end")
	return games, nil
}

// GetWithDetails retrieves a game with its frames and throws
func (r *gameRepository) GetWithDetails(c echo.Context, gameID int) (*models.Game, error) {
	logger.Debug("GetWithDetails start")
	game, err := models.Games(
		qm.Where("id = ?", gameID),
		qm.Where("deleted_flg = ?", false),
		qm.Load(models.GameRels.Frames),
		qm.Load(models.GameRels.Throws),
	).One(c.Request().Context(), r.con)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug("GetWithDetails end")
	return game, nil
}