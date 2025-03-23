package di

import (
	"legend_score/repositories"

	"go.uber.org/dig"
)

func provideRepository(c *dig.Container) {
	setProvide(c, repositories.NewPlayerRepository)
	setProvide(c, repositories.NewTournamentRepository)
	setProvide(c, repositories.NewTournamentManageRepository)
	setProvide(c, repositories.NewScheduleRepository)
	setProvide(c, repositories.NewScheduleDetailRepository)
	setProvide(c, repositories.NewTournamentScoreRepository)
	setProvide(c, repositories.NewTournamentScoreDetailRepository)
	setProvide(c, repositories.NewAdminUserRepository)
	setProvide(c, repositories.NewAdminUserTokenRepository)
	setProvide(c, repositories.NewTournamentLeanRepository)
}