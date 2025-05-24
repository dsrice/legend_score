package usecases

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	"legend_score/entities"
	"legend_score/entities/db"
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