package di

import (
	"legend_score/repositories"

	"go.uber.org/dig"
)

func provideRepository(c *dig.Container) {
	setProvide(c, repositories.NewUserRepository)
}