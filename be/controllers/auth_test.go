package controllers_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"legend_score/controllers"
	"legend_score/controllers/request"
	"legend_score/controllers/response"
	"legend_score/entities"
	"legend_score/usecases/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// mockValidator is a simple validator that always returns nil
type mockValidator struct{}

func (m *mockValidator) Validate(i interface{}) error {
	return nil
}

func TestAuthController_Login(t *testing.T) {
	// Setup
	e := echo.New()

	// Setup a mock validator that always returns nil
	e.Validator = &mockValidator{}

	// Create mock usecase
	mockAuthUseCase := new(mock.AuthUseCase)

	// Create controller with mock usecase
	authController := controllers.NewAuthController(mockAuthUseCase)

	// Test cases
	tests := []struct {
		name           string
		requestBody    request.LoginRequest
		setupMock      func()
		expectedStatus int
		expectedResult bool
	}{
		{
			name: "Success",
			requestBody: request.LoginRequest{
				LoginID:  "testuser",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup expectations for ValidateLogin
				mockAuthUseCase.On("ValidateLogin", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.LoginEntity) bool {
					return entity.LoginID == "testuser" && entity.Password == "Password123"
				})).Return(nil)

				// Setup expectations for Login
				token := "jwt-token-example"
				mockAuthUseCase.On("Login", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.LoginEntity) bool {
					return entity.LoginID == "testuser" && entity.Password == "Password123"
				})).Return(&token, nil)
			},
			expectedStatus: http.StatusOK,
			expectedResult: true,
		},
		{
			name: "Invalid Login",
			requestBody: request.LoginRequest{
				LoginID:  "testuser2",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup expectations for ValidateLogin to return an error
				mockAuthUseCase.On("ValidateLogin", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.LoginEntity) bool {
					return entity.LoginID == "testuser2" && entity.Password == "Password123"
				})).Return(assert.AnError)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: false,
		},
		{
			name: "Login Failed",
			requestBody: request.LoginRequest{
				LoginID:  "testuser3",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup expectations for ValidateLogin
				mockAuthUseCase.On("ValidateLogin", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.LoginEntity) bool {
					return entity.LoginID == "testuser3" && entity.Password == "Password123"
				})).Return(nil)

				// Setup expectations for Login to return an error
				mockAuthUseCase.On("Login", mocklib.Anything, mocklib.MatchedBy(func(entity *entities.LoginEntity) bool {
					return entity.LoginID == "testuser3" && entity.Password == "Password123"
				})).Return(nil, assert.AnError)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations
			tc.setupMock()

			// Create request - Convert LoginRequest to JSON
			jsonData, err := json.Marshal(tc.requestBody)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(jsonData)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Perform request
			err = authController.Login(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)

			// Parse response
			var response response.LoginResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, response.Result)

			// Verify mock expectations
			mockAuthUseCase.AssertExpectations(t)
		})
	}
}