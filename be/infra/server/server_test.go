package server_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/dig"
	"legend_score/controllers/ci"
	"legend_score/infra/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockAuthController is a mock implementation of the AuthController interface
type MockAuthController struct {
	mock.Mock
}

func (m *MockAuthController) Login(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// MockUserController is a mock implementation of the UserController interface
type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) CreateUser(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockUserController) GetUsers(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockUserController) GetUser(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

// MockGameController is a mock implementation of the GameController interface
type MockGameController struct {
	mock.Mock
}

func (m *MockGameController) GetGames(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockGameController) GetGamesByUserID(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockGameController) GetGameWithDetails(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockGameController) RegisterThrow(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestNewServer(t *testing.T) {
	// Create mock controllers
	mockAuthController := new(MockAuthController)
	mockUserController := new(MockUserController)
	mockGameController := new(MockGameController)

	// Create a new server
	s := server.NewServer(struct {
		dig.In
		Auth ci.AuthController
		User ci.UserController
		Game ci.GameController
	}{
		Auth: mockAuthController,
		User: mockUserController,
		Game: mockGameController,
	})

	// Assert that the server is not nil
	assert.NotNil(t, s)

	// Assert that the controllers are set correctly
	assert.Equal(t, mockAuthController, s.Auth)
	assert.Equal(t, mockUserController, s.User)
	assert.Equal(t, mockGameController, s.Game)
}

func TestCustomValidator_Validate(t *testing.T) {
	// Create a new echo instance
	e := echo.New()

	// Set up the validator for the echo instance
	e.Validator = &server.CustomValidator{Validator: validator.New()}

	// Create a new server
	s := server.NewServer(struct {
		dig.In
		Auth ci.AuthController
		User ci.UserController
		Game ci.GameController
	}{
		Auth: new(MockAuthController),
		User: new(MockUserController),
		Game: new(MockGameController),
	})

	// Start the server (this will initialize the validator)
	go func() {
		s.Start()
	}()

	// Create a test struct with validation tags
	type TestStruct struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}

	// Test cases
	testCases := []struct {
		name          string
		input         interface{}
		expectedError bool
	}{
		{
			name: "Valid Input",
			input: TestStruct{
				Name:  "Test User",
				Email: "test@example.com",
			},
			expectedError: false,
		},
		{
			name: "Missing Required Field",
			input: TestStruct{
				Name:  "",
				Email: "test@example.com",
			},
			expectedError: true,
		},
		{
			name: "Invalid Email",
			input: TestStruct{
				Name:  "Test User",
				Email: "invalid-email",
			},
			expectedError: true,
		},
	}

	// Create a request to test the validator
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test the validator with each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Bind the input to the context
			c.Set("input", tc.input)

			// Validate the input
			err := e.Validator.Validate(tc.input)

			// Assert that the error matches the expectation
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestServer_routing(t *testing.T) {
	// Create mock controllers
	mockAuthController := new(MockAuthController)
	mockUserController := new(MockUserController)
	mockGameController := new(MockGameController)

	// Set up expectations for the controllers
	mockAuthController.On("Login", mock.Anything).Return(nil)
	mockUserController.On("CreateUser", mock.Anything).Return(nil)
	mockUserController.On("GetUsers", mock.Anything).Return(nil)
	mockUserController.On("GetUser", mock.Anything).Return(nil)
	mockGameController.On("GetGames", mock.Anything).Return(nil)
	mockGameController.On("GetGamesByUserID", mock.Anything).Return(nil)
	mockGameController.On("GetGameWithDetails", mock.Anything).Return(nil)

	// Create a new server
	s := server.NewServer(struct {
		dig.In
		Auth ci.AuthController
		User ci.UserController
		Game ci.GameController
	}{
		Auth: mockAuthController,
		User: mockUserController,
		Game: mockGameController,
	})

	// Start the server (this will set up the routes)
	go func() {
		s.Start()
	}()

	// Create a new echo instance for testing
	e := echo.New()

	// Test the login route
	t.Run("Login Route", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the login handler
		err := mockAuthController.Login(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the login method was called
		mockAuthController.AssertCalled(t, "Login", c)
	})

	// Test the create user route
	t.Run("Create User Route", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the create user handler
		err := mockUserController.CreateUser(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the create user method was called
		mockUserController.AssertCalled(t, "CreateUser", c)
	})

	// Test the get users route
	t.Run("Get Users Route", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/user", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the get users handler
		err := mockUserController.GetUsers(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the get users method was called
		mockUserController.AssertCalled(t, "GetUsers", c)
	})

	// Test the get user route
	t.Run("Get User Route", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/user/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("user_id")
		c.SetParamValues("1")

		// Call the get user handler
		err := mockUserController.GetUser(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the get user method was called
		mockUserController.AssertCalled(t, "GetUser", c)
	})

	// Test the get games by user ID route
	t.Run("Get Games By User ID Route", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/user/1/game", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("user_id")
		c.SetParamValues("1")

		// Call the get games by user ID handler
		err := mockGameController.GetGamesByUserID(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the get games by user ID method was called
		mockGameController.AssertCalled(t, "GetGamesByUserID", c)
	})

	// Test the get games route
	t.Run("Get Games Route", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/game", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the get games handler
		err := mockGameController.GetGames(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the get games method was called
		mockGameController.AssertCalled(t, "GetGames", c)
	})

	// Test the get game with details route
	t.Run("Get Game With Details Route", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/game/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("game_id")
		c.SetParamValues("1")

		// Call the get game with details handler
		err := mockGameController.GetGameWithDetails(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the get game with details method was called
		mockGameController.AssertCalled(t, "GetGameWithDetails", c)
	})
}