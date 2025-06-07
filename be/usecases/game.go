package usecases

import (
	"errors"
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	"legend_score/entities"
	"legend_score/entities/db"
	"legend_score/infra/database/models"
	"legend_score/infra/logger"
	"legend_score/repositories/ri"
	"legend_score/usecases/ui"
)

type gameUseCase struct {
	game ri.GameRepository
}

func NewGameUseCase(game ri.GameRepository) ui.GameUseCase {
	return &gameUseCase{
		game: game,
	}
}

func (uc *gameUseCase) GetGames(c echo.Context, e *entities.GetGamesEntity) error {
	logger.Debug("GetGames start")

	// Get all games
	games, err := uc.game.GetAll(c)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}

	// Convert to entity
	gameEntities := make([]db.GameEntity, len(games))
	for i, game := range games {
		var gameEntity db.GameEntity
		gameEntity.SetEntity(game)
		gameEntities[i] = gameEntity
	}

	e.Games = gameEntities

	logger.Debug("GetGames end")
	return nil
}

func (uc *gameUseCase) GetGamesByUserID(c echo.Context, e *entities.GetGamesByUserIDEntity) error {
	logger.Debug("GetGamesByUserID start")

	// Get games by user ID
	games, err := uc.game.GetByUserID(c, e.UserID)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}

	// Convert to entity
	gameEntities := make([]db.GameEntity, len(games))
	for i, game := range games {
		var gameEntity db.GameEntity
		gameEntity.SetEntity(game)
		gameEntities[i] = gameEntity
	}

	e.Games = gameEntities

	logger.Debug("GetGamesByUserID end")
	return nil
}

func (uc *gameUseCase) GetGameWithDetails(c echo.Context, e *entities.GetGameWithDetailsEntity) error {
	logger.Debug("GetGameWithDetails start")

	// Get game with details
	game, err := uc.game.GetWithDetails(c, e.GameID)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return err
	}

	// Convert to entity
	var gameEntity db.GameEntity
	gameEntity.SetEntity(game)
	e.Game = gameEntity

	logger.Debug("GetGameWithDetails end")
	return nil
}

func (uc *gameUseCase) RegisterThrow(c echo.Context, e *entities.RegisterThrowEntity) error {
	logger.Debug("Start RegisterThrow")

	// Get user ID from context
	userID := c.Get("user_id").(int)

	// Validate frame and throw counts
	if e.FrameCount < 1 || e.FrameCount > 10 {
		e.Code = ecode.E0001
		return errors.New("invalid frame count")
	}

	if e.ThrowCount < 1 || e.ThrowCount > 2 {
		// Allow 3 throws in the 10th frame
		if !(e.FrameCount == 10 && e.ThrowCount == 3) {
			e.Code = ecode.E0001
			return errors.New("invalid throw count")
		}
	}

	// Create throw model
	throw := &models.Throw{
		UserID:     userID,
		GameID:     e.GameID,
		ThrowCount: e.ThrowCount,
		ThrowScore: e.ThrowScore,
		StrikeFlag: e.StrikeFlag,
		SpareFlag:  e.SpareFlag,
		Pin1:       e.Pin1,
		Pin2:       e.Pin2,
		Pin3:       e.Pin3,
		Pin4:       e.Pin4,
		Pin5:       e.Pin5,
		Pin6:       e.Pin6,
		Pin7:       e.Pin7,
		Pin8:       e.Pin8,
		Pin9:       e.Pin9,
		Pin10:      e.Pin10,
	}

	// Get or create frame
	frame, err := uc.game.GetFrameByGameIDAndFrameCount(c, e.GameID, e.FrameCount)
	if err != nil {
		e.Code = ecode.E9000
		return err
	}

	if frame == nil {
		// Create new frame if it doesn't exist
		// Create a new frame with default values
		frame = &models.Frame{
			UserID:     userID,
			GameID:     e.GameID,
		}
		frameID, err := uc.game.CreateFrame(c, frame)
		if err != nil {
			e.Code = ecode.E9000
			return err
		}
		throw.FrameID = frameID
	} else {
		throw.FrameID = frame.ID
	}

	// Register throw
	err = uc.game.RegisterThrow(c, throw)
	if err != nil {
		e.Code = ecode.E9000
		return err
	}

	logger.Debug("End RegisterThrow")
	return nil
}