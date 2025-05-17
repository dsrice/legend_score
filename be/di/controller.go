package di

import (
	"legend_score/controllers"

	"go.uber.org/dig"
)

func provideController(c *dig.Container) {
	setProvide(c, controllers.NewAuthController)
	setProvide(c, controllers.NewUserController)
}