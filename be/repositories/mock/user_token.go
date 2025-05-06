package mock

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"legend_score/infra/database/models"
	"legend_score/repositories/ri"
)

// UserTokenRepository is a mock implementation of ri.UserTokenRepository
type UserTokenRepository struct {
	mock.Mock
}

// Ensure UserTokenRepository implements ri.UserTokenRepository
var _ ri.UserTokenRepository = (*UserTokenRepository)(nil)

// Insert mocks the Insert method
func (m *UserTokenRepository) Insert(c echo.Context, ut *models.UserToken) error {
	args := m.Called(c, ut)
	return args.Error(0)
}