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

func TestGameUseCase_RegisterThrow(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)
	ctx.Set("user_id", 1) // Set user_id in context

	// Create mock repository
	mockGameRepo := new(mock.GameRepository)

	// Create usecase with mock repository
	gameUseCase := usecases.NewGameUseCase(mockGameRepo)

	// Test cases
	tests := []struct {
		name        string
		entity      *entities.RegisterThrowEntity
		setupMock   func()
		expectError bool
		expectCode  string
	}{
		{
			name: "Success - New Frame",
			entity: &entities.RegisterThrowEntity{
				GameID:     1,
				FrameCount: 1,
				ThrowCount: 1,
				ThrowScore: 6,
				StrikeFlag: false,
				SpareFlag:  false,
				Pin1:       1,
				Pin2:       1,
				Pin3:       1,
				Pin4:       0,
				Pin5:       1,
				Pin6:       0,
				Pin7:       1,
				Pin8:       0,
				Pin9:       1,
				Pin10:      0,
			},
			setupMock: func() {
				// Frame doesn't exist yet
				mockGameRepo.On("GetFrameByGameIDAndFrameCount", mocklib.Anything, 1, 1).Return(nil, nil)

				// Create new frame
				mockGameRepo.On("CreateFrame", mocklib.Anything, mocklib.MatchedBy(func(frame *models.Frame) bool {
					return frame.UserID == 1 && frame.GameID == 1
				})).Return(1, nil)

				// Register throw
				mockGameRepo.On("RegisterThrow", mocklib.Anything, mocklib.MatchedBy(func(throw *models.Throw) bool {
					return throw.UserID == 1 && throw.GameID == 1 && throw.FrameID == 1 && throw.ThrowCount == 1
				})).Return(nil)
			},
			expectError: false,
			expectCode:  "",
		},
		{
			name: "Success - Existing Frame",
			entity: &entities.RegisterThrowEntity{
				GameID:     1,
				FrameCount: 1,
				ThrowCount: 2,
				ThrowScore: 4,
				StrikeFlag: false,
				SpareFlag:  true,
				Pin1:       0,
				Pin2:       0,
				Pin3:       0,
				Pin4:       1,
				Pin5:       1,
				Pin6:       1,
				Pin7:       0,
				Pin8:       1,
				Pin9:       0,
				Pin10:      0,
			},
			setupMock: func() {
				// Frame exists
				frame := &models.Frame{
					ID:     1,
					UserID: 1,
					GameID: 1,
				}
				mockGameRepo.On("GetFrameByGameIDAndFrameCount", mocklib.Anything, 1, 1).Return(frame, nil)

				// Register throw
				mockGameRepo.On("RegisterThrow", mocklib.Anything, mocklib.MatchedBy(func(throw *models.Throw) bool {
					return throw.UserID == 1 && throw.GameID == 1 && throw.FrameID == 1 && throw.ThrowCount == 2
				})).Return(nil)
			},
			expectError: false,
			expectCode:  "",
		},
		{
			name: "Invalid Frame Count",
			entity: &entities.RegisterThrowEntity{
				GameID:     1,
				FrameCount: 11, // Invalid (> 10)
				ThrowCount: 1,
			},
			setupMock:   func() {},
			expectError: true,
			expectCode:  ecode.E0001,
		},
		{
			name: "Invalid Throw Count",
			entity: &entities.RegisterThrowEntity{
				GameID:     1,
				FrameCount: 1,
				ThrowCount: 3, // Invalid for non-10th frame
			},
			setupMock:   func() {},
			expectError: true,
			expectCode:  ecode.E0001,
		},
		{
			name: "GetFrameByGameIDAndFrameCount Error",
			entity: &entities.RegisterThrowEntity{
				GameID:     1,
				FrameCount: 1,
				ThrowCount: 1,
			},
			setupMock: func() {
				mockGameRepo.On("GetFrameByGameIDAndFrameCount", mocklib.Anything, 1, 1).Return(nil, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
		},
		{
			name: "CreateFrame Error",
			entity: &entities.RegisterThrowEntity{
				GameID:     1,
				FrameCount: 1,
				ThrowCount: 1,
			},
			setupMock: func() {
				mockGameRepo.On("GetFrameByGameIDAndFrameCount", mocklib.Anything, 1, 1).Return(nil, nil)
				mockGameRepo.On("CreateFrame", mocklib.Anything, mocklib.Anything).Return(0, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
		},
		{
			name: "RegisterThrow Error",
			entity: &entities.RegisterThrowEntity{
				GameID:     1,
				FrameCount: 1,
				ThrowCount: 1,
			},
			setupMock: func() {
				mockGameRepo.On("GetFrameByGameIDAndFrameCount", mocklib.Anything, 1, 1).Return(nil, nil)
				mockGameRepo.On("CreateFrame", mocklib.Anything, mocklib.Anything).Return(1, nil)
				mockGameRepo.On("RegisterThrow", mocklib.Anything, mocklib.Anything).Return(errors.New("database error"))
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

			// Call the method
			err := gameUseCase.RegisterThrow(ctx, tc.entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, tc.entity.Code)
			} else {
				assert.NoError(t, err)
			}

			// Verify mock expectations
			mockGameRepo.AssertExpectations(t)
		})
	}
}