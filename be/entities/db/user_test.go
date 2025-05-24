package db_test

import (
	"legend_score/entities/db"
	"legend_score/infra/database/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

func TestUserEntity_SetEntity(t *testing.T) {
	// Create a test user model
	errorTime := time.Now()
	lockTime := time.Now().Add(time.Hour)
	user := &models.User{
		ID:             1,
		LoginID:        "testuser",
		Name:           "Test User",
		Password:       "hashedpassword",
		ChangePassFlag: true,
		ErrorCount:     3,
		ErrorDatetime:  null.Time{Time: errorTime, Valid: true},
		LockDatetime:   null.Time{Time: lockTime, Valid: true},
	}

	// Create a UserEntity
	entity := &db.UserEntity{}

	// Call the method to test
	entity.SetEntity(user)

	// Assert the values were set correctly
	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, "testuser", entity.LoginID)
	assert.Equal(t, "Test User", entity.Name)
	assert.Equal(t, "hashedpassword", entity.Password)
	assert.True(t, entity.ChangePassFlag)
	assert.Equal(t, 3, entity.ErrorCount)
	assert.Equal(t, errorTime, *entity.ErrorDateTime)
	assert.Equal(t, lockTime, *entity.LockDateTime)
}

func TestUserEntity_SetEntity_NullTimes(t *testing.T) {
	// Create a test user model with null times
	user := &models.User{
		ID:             1,
		LoginID:        "testuser",
		Name:           "Test User",
		Password:       "hashedpassword",
		ChangePassFlag: true,
		ErrorCount:     3,
		ErrorDatetime:  null.Time{Valid: false},
		LockDatetime:   null.Time{Valid: false},
	}

	// Create a UserEntity
	entity := &db.UserEntity{}

	// Call the method to test
	entity.SetEntity(user)

	// Assert the values were set correctly
	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, "testuser", entity.LoginID)
	assert.Equal(t, "Test User", entity.Name)
	assert.Equal(t, "hashedpassword", entity.Password)
	assert.True(t, entity.ChangePassFlag)
	assert.Equal(t, 3, entity.ErrorCount)
	assert.Nil(t, entity.ErrorDateTime)
	assert.Nil(t, entity.LockDateTime)
}