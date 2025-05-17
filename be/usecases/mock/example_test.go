package mock_test

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"legend_score/entities"
	"legend_score/usecases/mock"
	"testing"
)

func TestAuthUseCaseMock(t *testing.T) {
	// Create a new mock instance
	authUseCase := new(mock.AuthUseCase)

	// Create test data
	loginEntity := &entities.LoginEntity{
		LoginID:  "testuser",
		Password: "password123",
	}
	token := "jwt-token-example"

	// Setup expectations
	authUseCase.On("ValidateLogin", mocklib.Anything, loginEntity).Return(nil)
	authUseCase.On("ValidatePassword", "password123").Return(true)
	authUseCase.On("ValidatePassword", "wrongpassword").Return(false)
	authUseCase.On("Login", mocklib.Anything, loginEntity).Return(&token, nil)

	// Create a context for testing
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Test ValidateLogin
	err := authUseCase.ValidateLogin(ctx, loginEntity)
	assert.NoError(t, err)

	// Test ValidatePassword with correct password
	result := authUseCase.ValidatePassword("password123")
	assert.True(t, result)

	// Test ValidatePassword with incorrect password
	result = authUseCase.ValidatePassword("wrongpassword")
	assert.False(t, result)

	// Test Login
	tokenResult, err := authUseCase.Login(ctx, loginEntity)
	assert.NoError(t, err)
	assert.Equal(t, &token, tokenResult)

	// Verify all expectations were met
	authUseCase.AssertExpectations(t)
}

func TestUserUseCaseMock(t *testing.T) {
	// Create a new mock instance
	userUseCase := new(mock.UserUseCase)

	// Create test data
	createUserEntity := &entities.CreateUserEntity{
		LoginID:  "newuser",
		Password: "password123",
		Name:     "New User",
	}
	getUsersEntity := &entities.GetUsersEntity{}
	getUserEntity := &entities.GetUserEntity{UserID: 1}

	// Setup expectations
	userUseCase.On("ValidateCreateUser", mocklib.Anything, createUserEntity).Return(nil)
	userUseCase.On("CreateUser", mocklib.Anything, createUserEntity).Return(nil)
	userUseCase.On("GetUsers", mocklib.Anything, getUsersEntity).Return(nil)
	userUseCase.On("GetUser", mocklib.Anything, getUserEntity).Return(nil)

	// Create a context for testing
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Test ValidateCreateUser
	err := userUseCase.ValidateCreateUser(ctx, createUserEntity)
	assert.NoError(t, err)

	// Test CreateUser
	err = userUseCase.CreateUser(ctx, createUserEntity)
	assert.NoError(t, err)

	// Test GetUsers
	err = userUseCase.GetUsers(ctx, getUsersEntity)
	assert.NoError(t, err)

	// Test GetUser
	err = userUseCase.GetUser(ctx, getUserEntity)
	assert.NoError(t, err)

	// Verify all expectations were met
	userUseCase.AssertExpectations(t)
}

func TestGameUseCaseMock(t *testing.T) {
	// Create a new mock instance
	gameUseCase := new(mock.GameUseCase)

	// Create test data
	gamesEntity := &entities.GamesEntity{
		Games: []entities.GameEntity{
			{ID: 1, UserID: 1},
		},
	}
	gameDetailEntity := &entities.GameDetailEntity{
		Game: entities.GameEntity{
			ID:     1,
			UserID: 1,
		},
		Frames: []entities.FrameEntity{},
		Throws: []entities.ThrowEntity{},
	}

	// Setup expectations
	gameUseCase.On("GetGamesByUserID", mocklib.Anything, 1).Return(gamesEntity, nil)
	gameUseCase.On("GetGamesByUserID", mocklib.Anything, 999).Return(nil, errors.New("not found"))
	gameUseCase.On("GetGameDetails", mocklib.Anything, 1, 1).Return(gameDetailEntity, nil)

	// Create a context for testing
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Test GetGamesByUserID with existing user
	games, err := gameUseCase.GetGamesByUserID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, gamesEntity, games)

	// Test GetGamesByUserID with non-existent user
	games, err = gameUseCase.GetGamesByUserID(ctx, 999)
	assert.Error(t, err)
	assert.Nil(t, games)

	// Test GetGameDetails
	gameDetail, err := gameUseCase.GetGameDetails(ctx, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, gameDetailEntity, gameDetail)

	// Verify all expectations were met
	gameUseCase.AssertExpectations(t)
}