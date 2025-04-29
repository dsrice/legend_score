package mock_test

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"legend_score/infra/database/models"
	repoMock "legend_score/repositories/mock"
	"testing"
)

func TestUserRepositoryMock(t *testing.T) {
	// Create a new mock instance
	userRepo := new(repoMock.UserRepository)
	
	// Create test data
	user := &models.User{ID: 1, LoginID: "testuser"}
	users := models.UserSlice{user}
	
	// Setup expectations
	userRepo.On("GetLoginID", mocklib.Anything, "testuser").Return(user, nil)
	userRepo.On("GetLoginID", mocklib.Anything, "nonexistent").Return(nil, errors.New("not found"))
	userRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(users, nil)
	userRepo.On("Insert", mocklib.Anything, user).Return(nil)
	
	// Create a context for testing
	e := echo.New()
	ctx := e.NewContext(nil, nil)
	
	// Test GetLoginID with existing user
	result, err := userRepo.GetLoginID(ctx, "testuser")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	
	// Test GetLoginID with non-existent user
	result, err = userRepo.GetLoginID(ctx, "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, result)
	
	// Test Get
	allUsers, err := userRepo.Get(ctx, nil)
	assert.NoError(t, err)
	assert.Equal(t, users, allUsers)
	
	// Test Insert
	err = userRepo.Insert(ctx, user)
	assert.NoError(t, err)
	
	// Verify all expectations were met
	userRepo.AssertExpectations(t)
}

func TestGameRepositoryMock(t *testing.T) {
	// Create a new mock instance
	gameRepo := new(repoMock.GameRepository)
	
	// Create test data
	game := &models.Game{ID: 1, UserID: 1}
	games := []*models.Game{game}
	
	// Setup expectations
	gameRepo.On("GetByUserID", mocklib.Anything, 1).Return(games, nil)
	gameRepo.On("GetWithDetails", mocklib.Anything, 1).Return(game, nil)
	
	// Create a context for testing
	e := echo.New()
	ctx := e.NewContext(nil, nil)
	
	// Test GetByUserID
	result, err := gameRepo.GetByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, games, result)
	
	// Test GetWithDetails
	gameDetails, err := gameRepo.GetWithDetails(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, game, gameDetails)
	
	// Verify all expectations were met
	gameRepo.AssertExpectations(t)
}

func TestUserTokenRepositoryMock(t *testing.T) {
	// Create a new mock instance
	tokenRepo := new(repoMock.UserTokenRepository)
	
	// Create test data
	token := &models.UserToken{ID: 1, UserID: 1, Token: "test-token"}
	
	// Setup expectations
	tokenRepo.On("Insert", mocklib.Anything, token).Return(nil)
	
	// Create a context for testing
	e := echo.New()
	ctx := e.NewContext(nil, nil)
	
	// Test Insert
	err := tokenRepo.Insert(ctx, token)
	assert.NoError(t, err)
	
	// Verify all expectations were met
	tokenRepo.AssertExpectations(t)
}