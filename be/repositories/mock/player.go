package mock

import (
	"legend_score/infra/database/models"

	"github.com/stretchr/testify/mock"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MockPlayerRepository struct {
	mock.Mock
}

func (m *MockPlayerRepository) GetPlayers(search []qm.QueryMod) (models.PlayerSlice, error) {
	ret := m.Called(search)
	return ret.Get(0).(models.PlayerSlice), ret.Error(1)
}