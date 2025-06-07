package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"legend_score/consts/ecode"
	"legend_score/controllers"
	"legend_score/controllers/response"
	"legend_score/entities"
	"legend_score/entities/db"
	"legend_score/infra/server"
	"legend_score/usecases/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestGameController_GetGames(t *testing.T) {
	// Setup
	e := echo.New()

	// Create mock usecase
	mockGameUseCase := new(mock.GameUseCase)

	// Create controller with mock usecase
	gameController := controllers.NewGameController(mockGameUseCase)

	// Test cases
	tests := []struct {
		name           string
		setupMock      func()
		expectedStatus int
		expectedResult bool
		expectedGames  int
	}{
		{
			name: "Success",
			setupMock: func() {
				// Setup expectations for GetGames
				mockGameUseCase.On("GetGames", mocklib.Anything, mocklib.Anything).Run(func(args mocklib.Arguments) {
					// Set games in the entity
					entity := args.Get(1).(*entities.GetGamesEntity)
					gameDate := time.Now()
					entity.Games = []db.GameEntity{
						{ID: 1, UserID: 1, Score: 150, GameDate: gameDate},
						{ID: 2, UserID: 2, Score: 200, GameDate: gameDate},
					}
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
			expectedGames:  2,
		},
		{
			name: "Error",
			setupMock: func() {
				// Setup expectations for GetGames to return an error
				mockGameUseCase.On("GetGames", mocklib.Anything, mocklib.Anything).Return(errors.New("test error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResult: false,
			expectedGames:  0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockGameUseCase.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/game", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Perform request
			err := gameController.GetGames(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.GetGamesResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)
			assert.Equal(t, tc.expectedGames, len(response.Games))

			// Verify mock expectations
			mockGameUseCase.AssertExpectations(t)
		})
	}
}

func TestGameController_GetGamesByUserID(t *testing.T) {
	// Setup
	e := echo.New()

	// Create mock usecase
	mockGameUseCase := new(mock.GameUseCase)

	// Create controller with mock usecase
	gameController := controllers.NewGameController(mockGameUseCase)

	// Test cases
	tests := []struct {
		name           string
		userID         int
		setupMock      func()
		expectedStatus int
		expectedResult bool
		expectedGames  int
	}{
		{
			name:   "Success",
			userID: 1,
			setupMock: func() {
				// Setup expectations for GetGamesByUserID
				mockGameUseCase.On("GetGamesByUserID", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetGamesByUserIDEntity) bool {
					return entity.UserID == 1
				})).Run(func(args mocklib.Arguments) {
					// Set games in the entity
					entity := args.Get(1).(*entities.GetGamesByUserIDEntity)
					gameDate := time.Now()
					entity.Games = []db.GameEntity{
						{ID: 1, UserID: 1, Score: 150, GameDate: gameDate},
						{ID: 2, UserID: 1, Score: 200, GameDate: gameDate},
					}
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
			expectedGames:  2,
		},
		{
			name:   "Error",
			userID: 1,
			setupMock: func() {
				// Setup expectations for GetGamesByUserID to return an error
				mockGameUseCase.On("GetGamesByUserID", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetGamesByUserIDEntity) bool {
					return entity.UserID == 1
				})).Return(errors.New("test error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResult: false,
			expectedGames:  0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockGameUseCase.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/user/"+strconv.Itoa(tc.userID)+"/game", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("user_id")
			c.SetParamValues(strconv.Itoa(tc.userID))

			// Perform request
			err := gameController.GetGamesByUserID(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.GetGamesResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)
			assert.Equal(t, tc.expectedGames, len(response.Games))

			// Verify mock expectations
			mockGameUseCase.AssertExpectations(t)
		})
	}
}

func TestGameController_GetGameWithDetails(t *testing.T) {
	// Setup
	e := echo.New()

	// Create mock usecase
	mockGameUseCase := new(mock.GameUseCase)

	// Create controller with mock usecase
	gameController := controllers.NewGameController(mockGameUseCase)

	// Test cases
	tests := []struct {
		name           string
		gameID         int
		setupMock      func()
		expectedStatus int
		expectedResult bool
	}{
		{
			name:   "Success",
			gameID: 1,
			setupMock: func() {
				// Setup expectations for GetGameWithDetails
				mockGameUseCase.On("GetGameWithDetails", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetGameWithDetailsEntity) bool {
					return entity.GameID == 1
				})).Run(func(args mocklib.Arguments) {
					// Set game in the entity
					entity := args.Get(1).(*entities.GetGameWithDetailsEntity)
					gameDate := time.Now()
					entity.Game = db.GameEntity{
						ID:       1,
						UserID:   1,
						Score:    150,
						GameDate: gameDate,
					}
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
		},
		{
			name:   "Error",
			gameID: 1,
			setupMock: func() {
				// Setup expectations for GetGameWithDetails to return an error
				mockGameUseCase.On("GetGameWithDetails", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetGameWithDetailsEntity) bool {
					return entity.GameID == 1
				})).Return(errors.New("test error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockGameUseCase.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/game/"+strconv.Itoa(tc.gameID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("game_id")
			c.SetParamValues(strconv.Itoa(tc.gameID))

			// Perform request
			err := gameController.GetGameWithDetails(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.GetGameWithDetailsResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)

			// Verify mock expectations
			mockGameUseCase.AssertExpectations(t)
		})
	}
}

func TestGameController_RegisterThrow(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New()}

	// Create mock usecase
	mockGameUseCase := new(mock.GameUseCase)

	// Create controller with mock usecase
	gameController := controllers.NewGameController(mockGameUseCase)

	// Test cases
	tests := []struct {
		name           string
		gameID         int
		requestBody    map[string]interface{}
		setupMock      func()
		expectedStatus int
		expectedResult bool
	}{
		{
			name:   "Success",
			gameID: 1,
			requestBody: map[string]interface{}{
				"frame_count": 1,
				"throw_count": 1,
				"pin_1":       1,
				"pin_2":       1,
				"pin_3":       1,
				"pin_4":       0,
				"pin_5":       1,
				"pin_6":       0,
				"pin_7":       1,
				"pin_8":       0,
				"pin_9":       1,
				"pin_10":      0,
			},
			setupMock: func() {
				// Setup expectations for RegisterThrow
				mockGameUseCase.On("RegisterThrow", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.RegisterThrowEntity) bool {
					return entity.GameID == 1 && entity.FrameCount == 1 && entity.ThrowCount == 1
				})).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
		},
		{
			name:   "Invalid Request",
			gameID: 1,
			requestBody: map[string]interface{}{
				"frame_count": "invalid", // Invalid type
				"throw_count": 1,
			},
			setupMock:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedResult: false,
		},
		{
			name:   "UseCase Error",
			gameID: 1,
			requestBody: map[string]interface{}{
				"frame_count": 1,
				"throw_count": 1,
				"pin_1":       1,
				"pin_2":       0,
				"pin_3":       0,
				"pin_4":       0,
				"pin_5":       0,
				"pin_6":       0,
				"pin_7":       0,
				"pin_8":       0,
				"pin_9":       0,
				"pin_10":      0,
			},
			setupMock: func() {
				// Setup expectations for RegisterThrow to return an error
				mockGameUseCase.On("RegisterThrow", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.RegisterThrowEntity) bool {
					return entity.GameID == 1 && entity.FrameCount == 1 && entity.ThrowCount == 1
				})).Run(func(args mocklib.Arguments) {
					entity := args.Get(1).(*entities.RegisterThrowEntity)
					entity.Code = ecode.E0001
				}).Return(errors.New("test error"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockGameUseCase.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Create request body
			jsonBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/game/"+strconv.Itoa(tc.gameID)+"/throw", bytes.NewReader(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("game_id")
			c.SetParamValues(strconv.Itoa(tc.gameID))

			// Perform request
			err := gameController.RegisterThrow(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.RegisterThrowResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)

			// Verify mock expectations
			mockGameUseCase.AssertExpectations(t)
		})
	}
}