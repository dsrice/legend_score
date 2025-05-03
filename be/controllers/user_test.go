package controllers_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"legend_score/controllers"
	"legend_score/controllers/response"
	"legend_score/entities"
	"legend_score/entities/db"
	"legend_score/usecases/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUserController_CreateUser(t *testing.T) {
	// Setup
	e := echo.New()

	// Create mock usecase
	mockUserUseCase := new(mock.UserUseCase)

	// Create controller with mock usecase
	userController := controllers.NewUserController(mockUserUseCase)

	// Test cases
	tests := []struct {
		name           string
		requestBody    string
		setupMock      func()
		expectedStatus int
		expectedResult bool
	}{
		{
			name: "Success",
			requestBody: `{
				"login_id": "newuser",
				"password": "Password123",
				"name": "New User"
			}`,
			setupMock: func() {
				// Setup expectations for ValidateCreateUser
				mockUserUseCase.On("ValidateCreateUser", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.CreateUserEntity) bool {
					return entity.LoginID == "newuser" &&
						entity.Password == "Password123" &&
						entity.Name == "New User"
				})).Return(nil)

				// Setup expectations for CreateUser
				mockUserUseCase.On("CreateUser", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.CreateUserEntity) bool {
					return entity.LoginID == "newuser" &&
						entity.Password == "Password123" &&
						entity.Name == "New User"
				})).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
		},
		{
			name: "Validation Error",
			requestBody: `{
				"login_id": "newuser",
				"password": "Password123",
				"name": "New User"
			}`,
			setupMock: func() {
				// Setup expectations for ValidateCreateUser to return an error
				mockUserUseCase.On("ValidateCreateUser", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.CreateUserEntity) bool {
					return entity.LoginID == "newuser" &&
						entity.Password == "Password123" &&
						entity.Name == "New User"
				})).Return(assert.AnError)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: false,
		},
		{
			name: "Creation Error",
			requestBody: `{
				"login_id": "newuser",
				"password": "Password123",
				"name": "New User"
			}`,
			setupMock: func() {
				// Setup expectations for ValidateCreateUser
				mockUserUseCase.On("ValidateCreateUser", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.CreateUserEntity) bool {
					return entity.LoginID == "newuser" &&
						entity.Password == "Password123" &&
						entity.Name == "New User"
				})).Return(nil)

				// Setup expectations for CreateUser to return an error
				mockUserUseCase.On("CreateUser", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.CreateUserEntity) bool {
					return entity.LoginID == "newuser" &&
						entity.Password == "Password123" &&
						entity.Name == "New User"
				})).Return(assert.AnError)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations
			tc.setupMock()

			// Create request
			req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(tc.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Perform request
			err := userController.CreateUser(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.CreateUserResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)

			// Verify mock expectations
			mockUserUseCase.AssertExpectations(t)
		})
	}
}

func TestUserController_GetUsers(t *testing.T) {
	// Setup
	e := echo.New()

	// Create mock usecase
	mockUserUseCase := new(mock.UserUseCase)

	// Create controller with mock usecase
	userController := controllers.NewUserController(mockUserUseCase)

	// Test cases
	tests := []struct {
		name           string
		queryParams    map[string]string
		setupMock      func()
		expectedStatus int
		expectedResult bool
		expectedUsers  int
	}{
		{
			name:        "Success - No Filters",
			queryParams: map[string]string{},
			setupMock: func() {
				// Setup expectations for GetUsers
				mockUserUseCase.On("GetUsers", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetUsersEntity) bool {
					return entity.UserID == nil && entity.LoginID == nil && entity.Name == nil
				})).Run(func(args mocklib.Arguments) {
					// Set users in the entity
					entity := args.Get(1).(*entities.GetUsersEntity)
					entity.Users = []db.UserEntity{
						{ID: 1, LoginID: "user1", Name: "User One"},
						{ID: 2, LoginID: "user2", Name: "User Two"},
					}
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
			expectedUsers:  2,
		},
		{
			name: "Success - With Filters",
			queryParams: map[string]string{
				"user_id":  "1",
				"login_id": "user1",
				"name":     "User",
			},
			setupMock: func() {
				// Setup expectations for GetUsers with filters
				mockUserUseCase.On("GetUsers", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetUsersEntity) bool {
					userID := 1
					loginID := "user1"
					name := "User"
					return entity.UserID != nil && *entity.UserID == userID &&
						entity.LoginID != nil && *entity.LoginID == loginID &&
						entity.Name != nil && *entity.Name == name
				})).Run(func(args mocklib.Arguments) {
					// Set users in the entity
					entity := args.Get(1).(*entities.GetUsersEntity)
					entity.Users = []db.UserEntity{
						{ID: 1, LoginID: "user1", Name: "User One"},
					}
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
			expectedUsers:  1,
		},
		{
			name:        "Error",
			queryParams: map[string]string{},
			setupMock: func() {
				// Setup expectations for GetUsers to return an error
				mockUserUseCase.On("GetUsers", mocklib.Anything, mocklib.Anything).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResult: false,
			expectedUsers:  0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations
			tc.setupMock()

			// Create request with query parameters
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			q := req.URL.Query()
			for key, value := range tc.queryParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Perform request
			err := userController.GetUsers(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.GetUsersResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)
			assert.Equal(t, tc.expectedUsers, len(response.Users))

			// Verify mock expectations
			mockUserUseCase.AssertExpectations(t)
		})
	}
}

func TestUserController_GetUser(t *testing.T) {
	// Setup
	e := echo.New()

	// Create mock usecase
	mockUserUseCase := new(mock.UserUseCase)

	// Create controller with mock usecase
	userController := controllers.NewUserController(mockUserUseCase)

	// Test cases
	tests := []struct {
		name           string
		userID         int
		setupMock      func()
		expectedStatus int
		expectedResult bool
	}{
		{
			name:   "Success",
			userID: 1,
			setupMock: func() {
				// Setup expectations for GetUser
				mockUserUseCase.On("GetUser", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetUserEntity) bool {
					return entity.UserID == 1
				})).Run(func(args mocklib.Arguments) {
					// Set user in the entity
					entity := args.Get(1).(*entities.GetUserEntity)
					entity.User = db.UserEntity{
						ID:      1,
						LoginID: "user1",
						Name:    "User One",
					}
				}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
		},
		{
			name:   "User Not Found",
			userID: 999,
			setupMock: func() {
				// Setup expectations for GetUser to return an error
				mockUserUseCase.On("GetUser", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.GetUserEntity) bool {
					return entity.UserID == 999
				})).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations
			tc.setupMock()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/user/"+strconv.Itoa(tc.userID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("user_id")
			c.SetParamValues(strconv.Itoa(tc.userID))

			// Perform request
			err := userController.GetUser(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.GetUserResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)

			// Verify mock expectations
			mockUserUseCase.AssertExpectations(t)
		})
	}
}