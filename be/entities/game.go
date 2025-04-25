package entities

import (
	"legend_score/infra/database/models"
	"time"
)

// GameEntity represents a single game
type GameEntity struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Name     string    `json:"name"`
	Score    int       `json:"score"`
	Count    int       `json:"count"`
	GameDate time.Time `json:"game_date"`
}

// GamesEntity represents a collection of games
type GamesEntity struct {
	Games []GameEntity `json:"games"`
}

// FrameEntity represents a single frame in a game
type FrameEntity struct {
	ID         int  `json:"id"`
	UserID     int  `json:"user_id"`
	GameID     int  `json:"game_id"`
	FrameCount int  `json:"frame_count"`
	FrameScore int  `json:"frame_score"`
	StrikeFlag bool `json:"strike_flag"`
	SpareFlag  bool `json:"spare_flag"`
}

// ThrowEntity represents a single throw in a frame
type ThrowEntity struct {
	ID         int  `json:"id"`
	UserID     int  `json:"user_id"`
	GameID     int  `json:"game_id"`
	FrameID    int  `json:"frame_id"`
	ThrowCount int  `json:"throw_count"`
	ThrowScore int  `json:"throw_score"`
	StrikeFlag bool `json:"strike_flag"`
	SpareFlag  bool `json:"spare_flag"`
	SplitFlag  bool `json:"split_flag"`
	Pin1       int  `json:"pin_1"`
	Pin2       int  `json:"pin_2"`
	Pin3       int  `json:"pin_3"`
	Pin4       int  `json:"pin_4"`
	Pin5       int  `json:"pin_5"`
	Pin6       int  `json:"pin_6"`
	Pin7       int  `json:"pin_7"`
	Pin8       int  `json:"pin_8"`
	Pin9       int  `json:"pin_9"`
	Pin10      int  `json:"pin_10"`
}

// GameDetailEntity represents a game with its frames and throws
type GameDetailEntity struct {
	Game   GameEntity    `json:"game"`
	Frames []FrameEntity `json:"frames"`
	Throws []ThrowEntity `json:"throws"`
}

// SetGameEntity sets the GameEntity from a models.Game
func (e *GameEntity) SetGameEntity(g *models.Game) {
	e.ID = g.ID
	e.UserID = g.UserID
	if g.Name.Valid {
		e.Name = g.Name.String
	}
	e.Score = g.Score
	if g.Count.Valid {
		e.Count = g.Count.Int
	}
	if g.GameDate.Valid {
		e.GameDate = g.GameDate.Time
	}
}

// SetFrameEntity sets the FrameEntity from a models.Frame
func (e *FrameEntity) SetFrameEntity(f *models.Frame) {
	e.ID = f.ID
	e.UserID = f.UserID
	e.GameID = f.GameID
	// Convert Decimal to float64, then to int
	frameCountFloat, _ := f.FrameCount.Float64()
	e.FrameCount = int(frameCountFloat)
	if f.FrameScore.Valid {
		e.FrameScore = f.FrameScore.Int
	}
	if f.StrikeFlag.Valid {
		e.StrikeFlag = f.StrikeFlag.Bool
	}
	if f.SpareFlag.Valid {
		e.SpareFlag = f.SpareFlag.Bool
	}
}

// SetThrowEntity sets the ThrowEntity from a models.Throw
func (e *ThrowEntity) SetThrowEntity(t *models.Throw) {
	e.ID = t.ID
	e.UserID = t.UserID
	e.GameID = t.GameID
	e.FrameID = t.FrameID
	e.ThrowCount = t.ThrowCount
	e.ThrowScore = t.ThrowScore
	e.StrikeFlag = t.StrikeFlag
	e.SpareFlag = t.SpareFlag
	e.SplitFlag = t.SplitFlag
	e.Pin1 = t.Pin1
	e.Pin2 = t.Pin2
	e.Pin3 = t.Pin3
	e.Pin4 = t.Pin4
	e.Pin5 = t.Pin5
	e.Pin6 = t.Pin6
	e.Pin7 = t.Pin7
	e.Pin8 = t.Pin8
	e.Pin9 = t.Pin9
	e.Pin10 = t.Pin10
}
