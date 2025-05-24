package entities_test

import (
	"legend_score/entities"
	"legend_score/infra/database/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

func TestGameEntity_SetGameEntity(t *testing.T) {
	// Create a test game model
	gameDate := time.Now()
	name := "Test Game"
	count := 5
	game := &models.Game{
		ID:      1,
		UserID:  2,
		Name:    null.String{String: name, Valid: true},
		Score:   300,
		Count:   null.Int{Int: count, Valid: true},
		GameDate: null.Time{Time: gameDate, Valid: true},
	}

	// Create a GameEntity
	entity := &entities.GameEntity{}

	// Call the method to test
	entity.SetGameEntity(game)

	// Assert the values were set correctly
	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, 2, entity.UserID)
	assert.Equal(t, name, entity.Name)
	assert.Equal(t, 300, entity.Score)
	assert.Equal(t, count, entity.Count)
	assert.Equal(t, gameDate, entity.GameDate)
}

func TestGameEntity_SetGameEntity_NullValues(t *testing.T) {
	// Create a test game model with null values
	game := &models.Game{
		ID:      1,
		UserID:  2,
		Name:    null.String{Valid: false},
		Score:   300,
		Count:   null.Int{Valid: false},
		GameDate: null.Time{Valid: false},
	}

	// Create a GameEntity
	entity := &entities.GameEntity{}

	// Call the method to test
	entity.SetGameEntity(game)

	// Assert the values were set correctly
	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, 2, entity.UserID)
	assert.Equal(t, "", entity.Name)
	assert.Equal(t, 300, entity.Score)
	assert.Equal(t, 0, entity.Count)
	assert.Equal(t, time.Time{}, entity.GameDate)
}

func TestThrowEntity_SetThrowEntity(t *testing.T) {
	// Create a test throw model
	throw := &models.Throw{
		ID:         1,
		UserID:     2,
		GameID:     3,
		FrameID:    4,
		ThrowCount: 5,
		ThrowScore: 10,
		StrikeFlag: true,
		SpareFlag:  false,
		SplitFlag:  true,
		Pin1:       1,
		Pin2:       1,
		Pin3:       1,
		Pin4:       1,
		Pin5:       1,
		Pin6:       1,
		Pin7:       1,
		Pin8:       1,
		Pin9:       1,
		Pin10:      1,
	}

	// Create a ThrowEntity
	entity := &entities.ThrowEntity{}

	// Call the method to test
	entity.SetThrowEntity(throw)

	// Assert the values were set correctly
	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, 2, entity.UserID)
	assert.Equal(t, 3, entity.GameID)
	assert.Equal(t, 4, entity.FrameID)
	assert.Equal(t, 5, entity.ThrowCount)
	assert.Equal(t, 10, entity.ThrowScore)
	assert.True(t, entity.StrikeFlag)
	assert.False(t, entity.SpareFlag)
	assert.True(t, entity.SplitFlag)
	assert.Equal(t, 1, entity.Pin1)
	assert.Equal(t, 1, entity.Pin2)
	assert.Equal(t, 1, entity.Pin3)
	assert.Equal(t, 1, entity.Pin4)
	assert.Equal(t, 1, entity.Pin5)
	assert.Equal(t, 1, entity.Pin6)
	assert.Equal(t, 1, entity.Pin7)
	assert.Equal(t, 1, entity.Pin8)
	assert.Equal(t, 1, entity.Pin9)
	assert.Equal(t, 1, entity.Pin10)
}