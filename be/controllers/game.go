package controllers

import (
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	ci "legend_score/controllers/ci"
	"legend_score/controllers/response"
	"legend_score/entities"
	"legend_score/infra/logger"
	"legend_score/usecases/ui"
	"net/http"
	"strconv"
)

type gameController struct {
	uc ui.GameUseCase
}

func NewGameController(uc ui.GameUseCase) ci.GameController {
	return &gameController{
		uc: uc,
	}
}

// GetGames godoc
// @Summary Get a list of all games
// @Description Get a list of all games
// @Tags game
// @Accept json
// @Produce json
// @Success 200 {object} response.GetGamesResponse
// @Failure 400 {object} response.GetGamesResponse
// @Router /game [get]
func (gc *gameController) GetGames(c echo.Context) error {
	logger.Debug("Start GetGames")

	entity := entities.GetGamesEntity{}

	err := gc.uc.GetGames(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, entity.Code)
	}

	games := make([]response.GameResponse, len(entity.Games))
	for i, game := range entity.Games {
		games[i] = response.GameResponse{
			ID:       game.ID,
			UserID:   game.UserID,
			GameDate: game.GameDate,
			Score:    game.Score,
		}
	}

	res := response.GetGamesResponse{
		Result: true,
		Games:  games,
	}

	logger.Debug("End GetGames")
	return c.JSON(http.StatusOK, res)
}

// GetGamesByUserID godoc
// @Summary Get a list of games for a specific user
// @Description Get a list of games for a specific user
// @Tags game
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} response.GetGamesResponse
// @Failure 400 {object} response.GetGamesResponse
// @Router /user/{user_id}/game [get]
func (gc *gameController) GetGamesByUserID(c echo.Context) error {
	logger.Debug("Start GetGamesByUserID")

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	entity := entities.GetGamesByUserIDEntity{
		UserID: userID,
	}

	err = gc.uc.GetGamesByUserID(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, entity.Code)
	}

	games := make([]response.GameResponse, len(entity.Games))
	for i, game := range entity.Games {
		games[i] = response.GameResponse{
			ID:       game.ID,
			UserID:   game.UserID,
			GameDate: game.GameDate,
			Score:    game.Score,
		}
	}

	res := response.GetGamesResponse{
		Result: true,
		Games:  games,
	}

	logger.Debug("End GetGamesByUserID")
	return c.JSON(http.StatusOK, res)
}

// GetGameWithDetails godoc
// @Summary Get a game with its frames and throws
// @Description Get a game with its frames and throws
// @Tags game
// @Accept json
// @Produce json
// @Param game_id path int true "Game ID"
// @Success 200 {object} response.GetGameWithDetailsResponse
// @Failure 400 {object} response.GetGameWithDetailsResponse
// @Router /game/{game_id} [get]
func (gc *gameController) GetGameWithDetails(c echo.Context) error {
	logger.Debug("Start GetGameWithDetails")

	gameID, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	entity := entities.GetGameWithDetailsEntity{
		GameID: gameID,
	}

	err = gc.uc.GetGameWithDetails(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, entity.Code)
	}

	game := response.GameDetailResponse{
		ID:       entity.Game.ID,
		UserID:   entity.Game.UserID,
		GameDate: entity.Game.GameDate,
		Score:    entity.Game.Score,
		Frames:   entity.Game.Frames,
		Throws:   entity.Game.Throws,
	}

	res := response.GetGameWithDetailsResponse{
		Result: true,
		Game:   game,
	}

	logger.Debug("End GetGameWithDetails")
	return c.JSON(http.StatusOK, res)
}