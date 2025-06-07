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

// RegisterThrow godoc
// @Summary Register a throw for a specific game
// @Description Register a throw with pin information for a specific game
// @Tags game
// @Accept json
// @Produce json
// @Param game_id path int true "Game ID"
// @Param throw body request.RegisterThrowRequest true "Throw information"
// @Success 200 {object} response.RegisterThrowResponse
// @Failure 400 {object} response.RegisterThrowResponse
// @Router /game/{game_id}/throw [post]
func (gc *gameController) RegisterThrow(c echo.Context) error {
	logger.Debug("Start RegisterThrow")

	req := new(request.RegisterThrowRequest)
	if err := c.Bind(req); err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}

	// Get game_id from path parameter
	gameID, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, ecode.E0001)
	}
	req.GameID = gameID

	// Calculate throw score based on pin information
	throwScore := 0
	if req.Pin1 > 0 { throwScore++ }
	if req.Pin2 > 0 { throwScore++ }
	if req.Pin3 > 0 { throwScore++ }
	if req.Pin4 > 0 { throwScore++ }
	if req.Pin5 > 0 { throwScore++ }
	if req.Pin6 > 0 { throwScore++ }
	if req.Pin7 > 0 { throwScore++ }
	if req.Pin8 > 0 { throwScore++ }
	if req.Pin9 > 0 { throwScore++ }
	if req.Pin10 > 0 { throwScore++ }

	// Determine if it's a strike or spare
	isStrike := req.ThrowCount == 1 && throwScore == 10
	isSpare := req.ThrowCount == 2 && throwScore + (10 - throwScore) == 10 // This is a simplification

	entity := entities.RegisterThrowEntity{
		GameID:     req.GameID,
		FrameCount: req.FrameCount,
		ThrowCount: req.ThrowCount,
		ThrowScore: throwScore,
		StrikeFlag: isStrike,
		SpareFlag:  isSpare,
		Pin1:       req.Pin1,
		Pin2:       req.Pin2,
		Pin3:       req.Pin3,
		Pin4:       req.Pin4,
		Pin5:       req.Pin5,
		Pin6:       req.Pin6,
		Pin7:       req.Pin7,
		Pin8:       req.Pin8,
		Pin9:       req.Pin9,
		Pin10:      req.Pin10,
	}

	err = gc.uc.RegisterThrow(c, &entity)
	if err != nil {
		logger.Error(err.Error())
		return ErrorResponse(c, entity.Code)
	}

	res := response.RegisterThrowResponse{
		Result: true,
	}

	logger.Debug("End RegisterThrow")
	return c.JSON(http.StatusOK, res)
}