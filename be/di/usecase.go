package di

import (
	"legend_score/usecases"

	"go.uber.org/dig"
)

func provideUseCase(c *dig.Container) {
	setProvide(c, usecases.NewPlayerUseCase)
	setProvide(c, usecases.NewTournamentUseCase)
	setProvide(c, usecases.NewTournamentScoreUseCase)
	setProvide(c, usecases.NewLineUseCase)
	setProvide(c, usecases.NewRankUseCase)
	setProvide(c, usecases.NewCyptographyUseCase)
	setProvide(c, usecases.NewAuthUseCase)
	setProvide(c, usecases.NewRoundRobinUseCase)
	setProvide(c, usecases.NewStepLadderUseCase)
}