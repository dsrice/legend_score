package entities_test

import (
	"legend_score/entities"
	"legend_score/entities/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserEntity(t *testing.T) {
	// Create a GetUserEntity
	entity := &entities.GetUserEntity{
		Code:   "200",
		UserID: 1,
		User: db.UserEntity{
			ID:      1,
			LoginID: "testuser",
			Name:    "Test User",
		},
	}

	// Assert the values are correct
	assert.Equal(t, "200", entity.Code)
	assert.Equal(t, 1, entity.UserID)
	assert.Equal(t, 1, entity.User.ID)
	assert.Equal(t, "testuser", entity.User.LoginID)
	assert.Equal(t, "Test User", entity.User.Name)
}