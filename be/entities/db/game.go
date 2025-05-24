package db

import (
	"legend_score/infra/database/models"
	"time"
)

type GameEntity struct {
	ID       int
	UserID   int
	GameDate time.Time
	Score    int
	Frames   []*models.Frame
	Throws   []*models.Throw
}

func (g *GameEntity) SetEntity(game *models.Game) {
	g.ID = game.ID
	g.UserID = game.UserID
	if game.GameDate.Valid {
		g.GameDate = game.GameDate.Time
	}
	g.Score = game.Score
	
	// Set frames and throws if they are loaded
	if game.R != nil {
		if game.R.Frames != nil {
			g.Frames = game.R.Frames
		}
		if game.R.Throws != nil {
			g.Throws = game.R.Throws
		}
	}
}