package di_test

import (
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"legend_score/entities"
	"legend_score/infra/database/models"
)

// Mock implementation of AuthUseCase
type mockAuthUseCase struct{}

func (m *mockAuthUseCase) ValidateLogin(c echo.Context, entity *entities.LoginEntity) error {
	return nil
}
func (m *mockAuthUseCase) ValidatePassword(password string) bool {
	return true
}
func (m *mockAuthUseCase) Login(c echo.Context, e *entities.LoginEntity) (*string, error) {
	return nil, nil
}

// Mock implementation of UserUseCase
type mockUserUseCase struct{}

func (m *mockUserUseCase) ValidateCreateUser(c echo.Context, e *entities.CreateUserEntity) error {
	return nil
}
func (m *mockUserUseCase) CreateUser(c echo.Context, e *entities.CreateUserEntity) error {
	return nil
}
func (m *mockUserUseCase) GetUsers(c echo.Context, e *entities.GetUsersEntity) error {
	return nil
}
func (m *mockUserUseCase) GetUser(c echo.Context, e *entities.GetUserEntity) error {
	return nil
}

// Mock implementation of GameUseCase
type mockGameUseCase struct{}

func (m *mockGameUseCase) GetGames(c echo.Context, e *entities.GetGamesEntity) error {
	return nil
}
func (m *mockGameUseCase) GetGamesByUserID(c echo.Context, e *entities.GetGamesByUserIDEntity) error {
	return nil
}
func (m *mockGameUseCase) GetGameWithDetails(c echo.Context, e *entities.GetGameWithDetailsEntity) error {
	return nil
}

// Mock implementations of repositories
type mockUserRepository struct{}

func (m *mockUserRepository) Get(c echo.Context, condition []qm.QueryMod) (models.UserSlice, error) {
	return nil, nil
}
func (m *mockUserRepository) GetLoginID(c echo.Context, loginID string) (*models.User, error) {
	return nil, nil
}
func (m *mockUserRepository) Insert(c echo.Context, ut *models.User) error {
	return nil
}

type mockUserTokenRepository struct{}

func (m *mockUserTokenRepository) Insert(c echo.Context, ut *models.UserToken) error {
	return nil
}

type mockGameRepository struct{}

func (m *mockGameRepository) GetAll(c echo.Context) ([]*models.Game, error) {
	return nil, nil
}
func (m *mockGameRepository) GetByUserID(c echo.Context, userID int) ([]*models.Game, error) {
	return nil, nil
}
func (m *mockGameRepository) GetWithDetails(c echo.Context, gameID int) (*models.Game, error) {
	return nil, nil
}