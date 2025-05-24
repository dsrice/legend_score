package entities_test

import (
	"legend_score/entities"
	"legend_score/entities/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsersEntity_SetFilters(t *testing.T) {
	// Create test filter values
	userID := 1
	loginID := "testuser"
	name := "Test User"

	// Create a GetUsersEntity
	entity := &entities.GetUsersEntity{}

	// Call the method to test
	entity.SetFilters(&userID, &loginID, &name)

	// Assert the values were set correctly
	assert.Equal(t, &userID, entity.UserID)
	assert.Equal(t, &loginID, entity.LoginID)
	assert.Equal(t, &name, entity.Name)
}

func TestGetUsersEntity(t *testing.T) {
	// Create a GetUsersEntity
	entity := &entities.GetUsersEntity{
		Code: "200",
		Users: []db.UserEntity{
			{
				ID:      1,
				LoginID: "user1",
				Name:    "User One",
			},
			{
				ID:      2,
				LoginID: "user2",
				Name:    "User Two",
			},
		},
	}

	// Assert the values are correct
	assert.Equal(t, "200", entity.Code)
	assert.Len(t, entity.Users, 2)
	assert.Equal(t, 1, entity.Users[0].ID)
	assert.Equal(t, "user1", entity.Users[0].LoginID)
	assert.Equal(t, "User One", entity.Users[0].Name)
	assert.Equal(t, 2, entity.Users[1].ID)
	assert.Equal(t, "user2", entity.Users[1].LoginID)
	assert.Equal(t, "User Two", entity.Users[1].Name)
}