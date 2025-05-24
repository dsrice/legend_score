package response

import (
	"legend_score/infra/database/models"
	"time"
)

// GameResponse represents a single game in the response
type GameResponse struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	GameDate time.Time `json:"game_date"`
	Score    int       `json:"score"`
}

// GetGamesResponse represents the response for the GetGames and GetGamesByUserID endpoints
type GetGamesResponse struct {
	Result bool           `json:"result"`
	Games  []GameResponse `json:"games"`
}

// GameDetailResponse represents a game with its frames and throws in the response
type GameDetailResponse struct {
	ID       int            `json:"id"`
	UserID   int            `json:"user_id"`
	GameDate time.Time      `json:"game_date"`
	Score    int            `json:"score"`
	Frames   []*models.Frame `json:"frames"`
	Throws   []*models.Throw `json:"throws"`
}

// GetGameWithDetailsResponse represents the response for the GetGameWithDetails endpoint
type GetGameWithDetailsResponse struct {
	Result bool               `json:"result"`
	Game   GameDetailResponse `json:"game"`
}