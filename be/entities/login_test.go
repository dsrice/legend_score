package entities_test

import (
	"legend_score/entities"
	"legend_score/entities/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginEntity(t *testing.T) {
	// Create a LoginEntity
	entity := &entities.LoginEntity{
		LoginID:  "testuser",
		Password: "password123",
		Code:     "200",
		User: db.UserEntity{
			ID:      1,
			LoginID: "testuser",
			Name:    "Test User",
		},
	}

	// Assert the values are correct
	assert.Equal(t, "testuser", entity.LoginID)
	assert.Equal(t, "password123", entity.Password)
	assert.Equal(t, "200", entity.Code)
	assert.Equal(t, 1, entity.User.ID)
	assert.Equal(t, "testuser", entity.User.LoginID)
	assert.Equal(t, "Test User", entity.User.Name)
}