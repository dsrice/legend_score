package di

import (
	"legend_score/usecases"

	"go.uber.org/dig"
)

func provideUseCase(c *dig.Container) {
	setProvide(c, usecases.NewAuthUseCase)
}