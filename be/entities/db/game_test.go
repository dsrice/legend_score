package db_test

import (
	"legend_score/entities/db"
	"legend_score/infra/database/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

func TestGameEntity_SetEntity(t *testing.T) {
	// Create a test game model
	gameDate := time.Now()
	game := &models.Game{
		ID:       1,
		UserID:   2,
		GameDate: null.Time{Time: gameDate, Valid: true},
		Score:    300,
	}

	// Create frames and throws for testing
	frames := []*models.Frame{
		{ID: 1, GameID: 1},
		{ID: 2, GameID: 1},
	}
	throws := []*models.Throw{
		{ID: 1, GameID: 1, FrameID: 1},
		{ID: 2, GameID: 1, FrameID: 1},
		{ID: 3, GameID: 1, FrameID: 2},
	}

	// Create a GameEntity
	entity := &db.GameEntity{}

	// Call the method to test
	entity.SetEntity(game)

	// Manually set frames and throws for testing
	entity.Frames = frames
	entity.Throws = throws

	// Assert the values were set correctly
	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, 2, entity.UserID)
	assert.Equal(t, gameDate, entity.GameDate)
	assert.Equal(t, 300, entity.Score)
	assert.Len(t, entity.Frames, 2)
	assert.Equal(t, 1, entity.Frames[0].ID)
	assert.Equal(t, 1, entity.Frames[0].GameID)
	assert.Equal(t, 2, entity.Frames[1].ID)
	assert.Equal(t, 1, entity.Frames[1].GameID)
	assert.Len(t, entity.Throws, 3)
	assert.Equal(t, 1, entity.Throws[0].ID)
	assert.Equal(t, 1, entity.Throws[0].GameID)
	assert.Equal(t, 1, entity.Throws[0].FrameID)
	assert.Equal(t, 2, entity.Throws[1].ID)
	assert.Equal(t, 1, entity.Throws[1].GameID)
	assert.Equal(t, 1, entity.Throws[1].FrameID)
	assert.Equal(t, 3, entity.Throws[2].ID)
	assert.Equal(t, 1, entity.Throws[2].GameID)
	assert.Equal(t, 2, entity.Throws[2].FrameID)
}