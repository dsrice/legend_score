package mock

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"legend_score/infra/database/models"
	"legend_score/repositories/ri"
)

// UserRepository is a mock implementation of ri.UserRepository
type UserRepository struct {
	mock.Mock
}

// Ensure UserRepository implements ri.UserRepository
var _ ri.UserRepository = (*UserRepository)(nil)

// Get mocks the Get method
func (m *UserRepository) Get(c echo.Context, condition []qm.QueryMod) (models.UserSlice, error) {
	args := m.Called(c, condition)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(models.UserSlice), args.Error(1)
}

// GetLoginID mocks the GetLoginID method
func (m *UserRepository) GetLoginID(c echo.Context, loginID string) (*models.User, error) {
	args := m.Called(c, loginID)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*models.User), args.Error(1)
}

// Insert mocks the Insert method
func (m *UserRepository) Insert(c echo.Context, ut *models.User) error {
	args := m.Called(c, ut)
	return args.Error(0)
}