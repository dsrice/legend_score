package usecases_test

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"github.com/volatiletech/null/v8"
	"legend_score/consts/ecode"
	"legend_score/entities"
	"legend_score/infra/database/models"
	"legend_score/repositories/mock"
	"legend_score/usecases"
	"testing"
	"time"
)

func TestGameUseCase_GetGames(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repository
	mockGameRepo := new(mock.GameRepository)

	// Create usecase with mock repository
	gameUseCase := usecases.NewGameUseCase(mockGameRepo)

	// Test cases
	tests := []struct {
		name        string
		setupMock   func()
		expectError bool
		expectCode  string
		expectGames int
	}{
		{
			name: "Success",
			setupMock: func() {
				// Setup mock for GetAll to return games
				gameDate := time.Now()
				games := []*models.Game{
					{ID: 1, UserID: 1, Score: 150, GameDate: null.TimeFrom(gameDate)},
					{ID: 2, UserID: 2, Score: 200, GameDate: null.TimeFrom(gameDate)},
				}
				mockGameRepo.On("GetAll", mocklib.Anything).Return(games, nil)
			},
			expectError: false,
			expectCode:  "",
			expectGames: 2,
		},
		{
			name: "Repository Error",
			setupMock: func() {
				// Setup mock for GetAll to return error
				mockGameRepo.On("GetAll", mocklib.Anything).Return(nil, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
			expectGames: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockGameRepo.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Create entity
			entity := entities.GetGamesEntity{}

			// Call the method
			err := gameUseCase.GetGames(ctx, &entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, entity.Code)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectGames, len(entity.Games))
			}

			// Verify mock expectations
			mockGameRepo.AssertExpectations(t)
		})
	}
}

func TestGameUseCase_GetGamesByUserID(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repository
	mockGameRepo := new(mock.GameRepository)

	// Create usecase with mock repository
	gameUseCase := usecases.NewGameUseCase(mockGameRepo)

	// Test cases
	tests := []struct {
		name        string
		userID      int
		setupMock   func()
		expectError bool
		expectCode  string
		expectGames int
	}{
		{
			name:   "Success",
			userID: 1,
			setupMock: func() {
				// Setup mock for GetByUserID to return games
				gameDate := time.Now()
				games := []*models.Game{
					{ID: 1, UserID: 1, Score: 150, GameDate: null.TimeFrom(gameDate)},
					{ID: 2, UserID: 1, Score: 200, GameDate: null.TimeFrom(gameDate)},
				}
				mockGameRepo.On("GetByUserID", mocklib.Anything, 1).Return(games, nil)
			},
			expectError: false,
			expectCode:  "",
			expectGames: 2,
		},
		{
			name:   "Repository Error",
			userID: 1,
			setupMock: func() {
				// Setup mock for GetByUserID to return error
				mockGameRepo.On("GetByUserID", mocklib.Anything, 1).Return(nil, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
			expectGames: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockGameRepo.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Create entity
			entity := entities.GetGamesByUserIDEntity{
				UserID: tc.userID,
			}

			// Call the method
			err := gameUseCase.GetGamesByUserID(ctx, &entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, entity.Code)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectGames, len(entity.Games))
			}

			// Verify mock expectations
			mockGameRepo.AssertExpectations(t)
		})
	}
}

func TestGameUseCase_GetGameWithDetails(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repository
	mockGameRepo := new(mock.GameRepository)

	// Create usecase with mock repository
	gameUseCase := usecases.NewGameUseCase(mockGameRepo)

	// Test cases
	tests := []struct {
		name        string
		gameID      int
		setupMock   func()
		expectError bool
		expectCode  string
	}{
		{
			name:   "Success",
			gameID: 1,
			setupMock: func() {
				// Setup mock for GetWithDetails to return game
				gameDate := time.Now()
				game := &models.Game{
					ID:       1,
					UserID:   1,
					Score:    150,
					GameDate: null.TimeFrom(gameDate),
				}
				mockGameRepo.On("GetWithDetails", mocklib.Anything, 1).Return(game, nil)
			},
			expectError: false,
			expectCode:  "",
		},
		{
			name:   "Repository Error",
			gameID: 1,
			setupMock: func() {
				// Setup mock for GetWithDetails to return error
				mockGameRepo.On("GetWithDetails", mocklib.Anything, 1).Return(nil, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockGameRepo.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Create entity
			entity := entities.GetGameWithDetailsEntity{
				GameID: tc.gameID,
			}

			// Call the method
			err := gameUseCase.GetGameWithDetails(ctx, &entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, entity.Code)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.gameID, entity.Game.ID)
			}

			// Verify mock expectations
			mockGameRepo.AssertExpectations(t)
		})
	}
}