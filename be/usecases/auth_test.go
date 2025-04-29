package usecases_test

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mocklib "github.com/stretchr/testify/mock"
	"github.com/volatiletech/null/v8"
	"legend_score/consts/ecode"
	"legend_score/entities"
	"legend_score/entities/db"
	"legend_score/infra/database/models"
	"legend_score/repositories/mock"
	"legend_score/usecases"
	"testing"
	"time"
)

func TestAuthUseCase_ValidateLogin(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repositories
	mockUserRepo := new(mock.UserRepository)
	mockUserTokenRepo := new(mock.UserTokenRepository)

	// Create usecase with mock repositories
	authUseCase := usecases.NewAuthUseCase(mockUserRepo, mockUserTokenRepo)

	// Test cases
	tests := []struct {
		name       string
		loginID    string
		password   string
		setupMock  func()
		expectError bool
	}{
		{
			name:     "Success",
			loginID:  "testuser",
			password: "Password123",
			setupMock: func() {
				// Setup mock for GetLoginID
				user := &models.User{
					ID:           1,
					LoginID:      "testuser",
					Password:     "hashedpassword",
					LockDatetime: null.Time{},
				}
				mockUserRepo.On("GetLoginID", mocklib.Anything, "testuser").Return(user, nil)
			},
			expectError: false,
		},
		{
			name:     "Invalid Password",
			loginID:  "testuser",
			password: "weak",
			setupMock: func() {
				// No need to setup mock for repository as validation fails before repository call
			},
			expectError: true,
		},
		{
			name:     "User Not Found",
			loginID:  "nonexistent",
			password: "Password123",
			setupMock: func() {
				// Setup mock for GetLoginID to return error
				mockUserRepo.On("GetLoginID", mocklib.Anything, "nonexistent").Return(nil, errors.New("user not found"))
			},
			expectError: true,
		},
		{
			name:     "Account Locked",
			loginID:  "lockeduser",
			password: "Password123",
			setupMock: func() {
				// Setup mock for GetLoginID with locked account
				lockTime := time.Now() // Current time means account is locked
				user := &models.User{
					ID:           2,
					LoginID:      "lockeduser",
					Password:     "hashedpassword",
					LockDatetime: null.Time{Valid: true, Time: lockTime},
				}
				mockUserRepo.On("GetLoginID", mocklib.Anything, "lockeduser").Return(user, nil)
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations
			tc.setupMock()

			// Create login entity
			entity := &entities.LoginEntity{
				LoginID:  tc.loginID,
				Password: tc.password,
			}

			// Call the method
			err := authUseCase.ValidateLogin(ctx, entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify mock expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestAuthUseCase_ValidatePassword(t *testing.T) {
	// Setup
	mockUserRepo := new(mock.UserRepository)
	mockUserTokenRepo := new(mock.UserTokenRepository)

	// Create usecase with mock repositories
	authUseCase := usecases.NewAuthUseCase(mockUserRepo, mockUserTokenRepo)

	// Test cases
	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "Valid Password",
			password: "Password123",
			expected: true,
		},
		{
			name:     "Too Short",
			password: "Pass1",
			expected: false,
		},
		{
			name:     "No Uppercase",
			password: "password123",
			expected: false,
		},
		{
			name:     "No Lowercase",
			password: "PASSWORD123",
			expected: false,
		},
		{
			name:     "No Numbers",
			password: "PasswordABC",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the method
			result := authUseCase.ValidatePassword(tc.password)

			// Assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestAuthUseCase_Login(t *testing.T) {
	// Setup
	e := echo.New()
	ctx := e.NewContext(nil, nil)

	// Create mock repositories
	mockUserRepo := new(mock.UserRepository)
	mockUserTokenRepo := new(mock.UserTokenRepository)

	// Create usecase with mock repositories
	authUseCase := usecases.NewAuthUseCase(mockUserRepo, mockUserTokenRepo)

	// Test cases
	tests := []struct {
		name       string
		setupEntity func() *entities.LoginEntity
		setupMock  func()
		expectToken bool
		expectError bool
	}{
		{
			name: "Success",
			setupEntity: func() *entities.LoginEntity {
				// This is a simplified test - in reality, the password would be hashed
				// and compared with the stored hash, but for testing we're using a direct match
				return &entities.LoginEntity{
					LoginID:  "testuser",
					Password: "Password123",
					User: db.UserEntity{
						ID:       1,
						LoginID:  "testuser",
						Password: "legend_score_salt_dev", // This would be a hash in reality
					},
				}
			},
			setupMock: func() {
				// Setup mock for Insert token
				mockUserTokenRepo.On("Insert", mocklib.Anything, mocklib.AnythingOfType("*models.UserToken")).Return(nil)
			},
			expectToken: true,
			expectError: false,
		},
		{
			name: "Token Creation Error",
			setupEntity: func() *entities.LoginEntity {
				return &entities.LoginEntity{
					LoginID:  "testuser",
					Password: "Password123",
					User: db.UserEntity{
						ID:       1,
						LoginID:  "testuser",
						Password: "legend_score_salt_dev", // This would be a hash in reality
					},
				}
			},
			setupMock: func() {
				// Setup mock for Insert token to return error
				mockUserTokenRepo.On("Insert", mocklib.Anything, mocklib.AnythingOfType("*models.UserToken")).Return(errors.New("token creation error"))
			},
			expectToken: false,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations
			tc.setupMock()

			// Create login entity
			entity := tc.setupEntity()

			// Call the method
			token, err := authUseCase.Login(ctx, entity)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, token)
				assert.Equal(t, ecode.E0001, entity.Code)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, token)
				assert.NotEmpty(t, *token)
			}

			// Verify mock expectations
			mockUserTokenRepo.AssertExpectations(t)
		})
	}
}