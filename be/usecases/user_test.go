package usecases_test

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"legend_score/consts/ecode"
	"legend_score/entities"
	"legend_score/infra/database/models"
	"legend_score/repositories/mock"
	"legend_score/usecases"
	usecaseMock "legend_score/usecases/mock"
	"testing"
)

func TestUserUseCase_ValidateCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repositories
	mockUserRepo := new(mock.UserRepository)
	mockAuthUseCase := new(usecaseMock.AuthUseCase)

	// Create usecase with mock repositories
	userUseCase := usecases.NewUserUseCase(mockUserRepo, mockAuthUseCase)

	// Test cases
	tests := []struct {
		name        string
		entity      *entities.CreateUserEntity
		setupMock   func()
		expectError bool
		expectCode  string
	}{
		{
			name: "Success",
			entity: &entities.CreateUserEntity{
				LoginID:  "newuser",
				Name:     "New User",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup mock for Get to return empty slice (no existing user)
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(models.UserSlice{}, nil)
				// Setup mock for ValidatePassword to return true
				mockAuthUseCase.On("ValidatePassword", "Password123").Return(true)
			},
			expectError: false,
			expectCode:  "",
		},
		{
			name: "LoginID Already Used",
			entity: &entities.CreateUserEntity{
				LoginID:  "existinguser",
				Name:     "Existing User",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup mock for Get to return a user (login ID already exists)
				existingUser := models.User{
					ID:      1,
					LoginID: "existinguser",
					Name:    "Existing User",
				}
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(models.UserSlice{&existingUser}, nil)
			},
			expectError: true,
			expectCode:  ecode.E2001,
		},
		{
			name: "Invalid Password",
			entity: &entities.CreateUserEntity{
				LoginID:  "newuser",
				Name:     "New User",
				Password: "weak",
			},
			setupMock: func() {
				// Setup mock for Get to return empty slice (no existing user)
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(models.UserSlice{}, nil)
				// Setup mock for ValidatePassword to return false
				mockAuthUseCase.On("ValidatePassword", "weak").Return(false)
			},
			expectError: true,
			expectCode:  ecode.E2002,
		},
		{
			name: "Repository Error",
			entity: &entities.CreateUserEntity{
				LoginID:  "newuser",
				Name:     "New User",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup mock for Get to return error
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(nil, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockUserRepo.ExpectedCalls = nil
			mockAuthUseCase.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Call the method
			err := userUseCase.ValidateCreateUser(ctx, tc.entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, tc.entity.Code)
			} else {
				assert.NoError(t, err)
			}

			// Verify mock expectations
			mockUserRepo.AssertExpectations(t)
			mockAuthUseCase.AssertExpectations(t)
		})
	}
}

func TestUserUseCase_CreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repositories
	mockUserRepo := new(mock.UserRepository)
	mockAuthUseCase := new(usecaseMock.AuthUseCase)

	// Create usecase with mock repositories
	userUseCase := usecases.NewUserUseCase(mockUserRepo, mockAuthUseCase)

	// Test cases
	tests := []struct {
		name        string
		entity      *entities.CreateUserEntity
		setupMock   func()
		expectError bool
		expectCode  string
	}{
		{
			name: "Success",
			entity: &entities.CreateUserEntity{
				LoginID:  "newuser",
				Name:     "New User",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup mock for Insert to return nil (success)
				mockUserRepo.On("Insert", mocklib.Anything, mocklib.AnythingOfType("*models.User")).Return(nil)
			},
			expectError: false,
			expectCode:  "",
		},
		{
			name: "Insert Error",
			entity: &entities.CreateUserEntity{
				LoginID:  "newuser",
				Name:     "New User",
				Password: "Password123",
			},
			setupMock: func() {
				// Setup mock for Insert to return error
				mockUserRepo.On("Insert", mocklib.Anything, mocklib.AnythingOfType("*models.User")).Return(errors.New("insert error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockUserRepo.ExpectedCalls = nil
			mockAuthUseCase.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Call the method
			err := userUseCase.CreateUser(ctx, tc.entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, tc.entity.Code)
			} else {
				assert.NoError(t, err)
			}

			// Verify mock expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestUserUseCase_GetUsers(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repositories
	mockUserRepo := new(mock.UserRepository)
	mockAuthUseCase := new(usecaseMock.AuthUseCase)

	// Create usecase with mock repositories
	userUseCase := usecases.NewUserUseCase(mockUserRepo, mockAuthUseCase)

	// Test data
	userID := 1
	loginID := "testuser"
	name := "Test User"

	// Test cases
	tests := []struct {
		name        string
		entity      *entities.GetUsersEntity
		setupMock   func()
		expectError bool
		expectCode  string
		expectUsers int
	}{
		{
			name: "Success - No Filters",
			entity: &entities.GetUsersEntity{},
			setupMock: func() {
				// Setup mock for Get to return users
				users := models.UserSlice{
					&models.User{ID: 1, LoginID: "user1", Name: "User One"},
					&models.User{ID: 2, LoginID: "user2", Name: "User Two"},
				}
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(users, nil)
			},
			expectError: false,
			expectCode:  "",
			expectUsers: 2,
		},
		{
			name: "Success - With UserID Filter",
			entity: &entities.GetUsersEntity{
				UserID: &userID,
			},
			setupMock: func() {
				// Setup mock for Get to return filtered users
				users := models.UserSlice{
					&models.User{ID: 1, LoginID: "user1", Name: "User One"},
				}
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(users, nil)
			},
			expectError: false,
			expectCode:  "",
			expectUsers: 1,
		},
		{
			name: "Success - With LoginID Filter",
			entity: &entities.GetUsersEntity{
				LoginID: &loginID,
			},
			setupMock: func() {
				// Setup mock for Get to return filtered users
				users := models.UserSlice{
					&models.User{ID: 1, LoginID: "testuser", Name: "Test User"},
				}
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(users, nil)
			},
			expectError: false,
			expectCode:  "",
			expectUsers: 1,
		},
		{
			name: "Success - With Name Filter",
			entity: &entities.GetUsersEntity{
				Name: &name,
			},
			setupMock: func() {
				// Setup mock for Get to return filtered users
				users := models.UserSlice{
					&models.User{ID: 1, LoginID: "testuser", Name: "Test User"},
				}
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(users, nil)
			},
			expectError: false,
			expectCode:  "",
			expectUsers: 1,
		},
		{
			name: "Repository Error",
			entity: &entities.GetUsersEntity{},
			setupMock: func() {
				// Setup mock for Get to return error
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(nil, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
			expectUsers: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockUserRepo.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Call the method
			err := userUseCase.GetUsers(ctx, tc.entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, tc.entity.Code)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectUsers, len(tc.entity.Users))
			}

			// Verify mock expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestUserUseCase_GetUser(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repositories
	mockUserRepo := new(mock.UserRepository)
	mockAuthUseCase := new(usecaseMock.AuthUseCase)

	// Create usecase with mock repositories
	userUseCase := usecases.NewUserUseCase(mockUserRepo, mockAuthUseCase)

	// Test cases
	tests := []struct {
		name        string
		entity      *entities.GetUserEntity
		setupMock   func()
		expectError bool
		expectCode  string
	}{
		{
			name: "Success",
			entity: &entities.GetUserEntity{
				UserID: 1,
			},
			setupMock: func() {
				// Setup mock for Get to return user
				users := models.UserSlice{
					&models.User{ID: 1, LoginID: "user1", Name: "User One"},
				}
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(users, nil)
			},
			expectError: false,
			expectCode:  "",
		},
		{
			name: "User Not Found",
			entity: &entities.GetUserEntity{
				UserID: 999,
			},
			setupMock: func() {
				// Setup mock for Get to return empty slice
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(models.UserSlice{}, nil)
			},
			expectError: true,
			expectCode:  ecode.E0001,
		},
		{
			name: "Repository Error",
			entity: &entities.GetUserEntity{
				UserID: 1,
			},
			setupMock: func() {
				// Setup mock for Get to return error
				mockUserRepo.On("Get", mocklib.Anything, mocklib.Anything).Return(nil, errors.New("database error"))
			},
			expectError: true,
			expectCode:  ecode.E9000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockUserRepo.ExpectedCalls = nil

			// Setup mock expectations
			tc.setupMock()

			// Call the method
			err := userUseCase.GetUser(ctx, tc.entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectCode, tc.entity.Code)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, 1, tc.entity.User.ID)
			}

			// Verify mock expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}