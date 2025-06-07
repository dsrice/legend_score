package entities_test

import (
	"legend_score/controllers/request"
	"legend_score/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserEntity_SetEntity(t *testing.T) {
	// Create a test request
	req := &request.CreateUserRequest{
		LoginID:  "testuser",
		Name:     "Test User",
		Password: "password123",
	}

	// Create a CreateUserEntity
	entity := &entities.CreateUserEntity{}

	// Call the method to test
	entity.SetEntity(req)

	// Assert the values were set correctly
	assert.Equal(t, "testuser", entity.LoginID)
	assert.Equal(t, "Test User", entity.Name)
	assert.Equal(t, "password123", entity.Password)
}