package repositories

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"legend_score/infra/database/connection"
	"legend_score/infra/database/models"
	"legend_score/infra/logger"
	"legend_score/repositories/ri"
	"time"
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

// GetAll retrieves all games
func (r *gameRepository) GetAll(c echo.Context) ([]*models.Game, error) {
	logger.Debug("GetAll start")
	games, err := models.Games(
		qm.Where("deleted_flg = ?", false),
		qm.OrderBy("game_date DESC"),
	).All(c.Request().Context(), r.con)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug("GetAll end")
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

// GetFrameByGameIDAndFrameCount retrieves a frame by game ID and frame count
func (r *gameRepository) GetFrameByGameIDAndFrameCount(c echo.Context, gameID, frameCount int) (*models.Frame, error) {
	logger.Debug("GetFrameByGameIDAndFrameCount start")

	frame, err := models.Frames(
		qm.Where("game_id = ?", gameID),
		qm.Where("frame_count = ?", frameCount),
		qm.Where("deleted_flg = ?", false),
	).One(c.Request().Context(), r.con)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	logger.Debug("GetFrameByGameIDAndFrameCount end")
	return frame, nil
}

// CreateFrame creates a new frame
func (r *gameRepository) CreateFrame(c echo.Context, frame *models.Frame) (int, error) {
	logger.Debug("CreateFrame start")

	// Use ExecContext instead of Insert to avoid the SELECT query
	query := "INSERT INTO frames (user_id, game_id, frame_count, created_at, updated_at, deleted_flg) VALUES (?, ?, ?, ?, ?, ?)"

	// Set default values
	if frame.CreatedAt.IsZero() {
		frame.CreatedAt = time.Now()
	}
	if frame.UpdatedAt.IsZero() {
		frame.UpdatedAt = time.Now()
	}

	// Execute the query
	result, err := r.con.ExecContext(c.Request().Context(), query, 
		frame.UserID, frame.GameID, frame.FrameCount, frame.CreatedAt, frame.UpdatedAt, false)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	// Set the ID in the frame
	frame.ID = int(id)

	logger.Debug("CreateFrame end")
	return frame.ID, nil
}

// RegisterThrow registers a throw
func (r *gameRepository) RegisterThrow(c echo.Context, throw *models.Throw) error {
	logger.Debug("RegisterThrow start")

	// Use ExecContext instead of Insert to avoid the SELECT query
	query := "INSERT INTO throws (user_id, game_id, frame_id, throw_count, throw_score, strike_flag, spare_flag, split_flag, pin_1, pin_2, pin_3, pin_4, pin_5, pin_6, pin_7, pin_8, pin_9, pin_10, created_at, updated_at, deleted_flg) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Set default values
	if throw.CreatedAt.IsZero() {
		throw.CreatedAt = time.Now()
	}
	if throw.UpdatedAt.IsZero() {
		throw.UpdatedAt = time.Now()
	}

	// Execute the query
	result, err := r.con.ExecContext(c.Request().Context(), query, 
		throw.UserID, throw.GameID, throw.FrameID, throw.ThrowCount, throw.ThrowScore, 
		throw.StrikeFlag, throw.SpareFlag, throw.SplitFlag, 
		throw.Pin1, throw.Pin2, throw.Pin3, throw.Pin4, throw.Pin5, 
		throw.Pin6, throw.Pin7, throw.Pin8, throw.Pin9, throw.Pin10, 
		throw.CreatedAt, throw.UpdatedAt, false)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	// Set the ID in the throw
	throw.ID = int(id)

	logger.Debug("RegisterThrow end")
	return nil
}