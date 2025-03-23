package di

import (
	"legend_score/controllers"

	"go.uber.org/dig"
)

func provideController(c *dig.Container) {
	setProvide(c, controllers.NewAuthController)
	setProvide(c, controllers.NewLineController)
	setProvide(c, controllers.NewTournamentController)
	setProvide(c, controllers.NewTournamentManageController)
	setProvide(c, controllers.NewTournamentPlayerController)
	setProvide(c, controllers.NewTournamentScoreController)
	setProvide(c, controllers.NewRankController)
	setProvide(c, controllers.NewPlayerScoreController)
	setProvide(c, controllers.NewPlayerController)
}