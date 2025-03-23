package di

import (
	"legend_score/infra/database/connection"
	"legend_score/infra/logger"
	"legend_score/infra/server"

	"go.uber.org/dig"
)

// BuildContainer
// DI注入メイン処理
func BuildContainer(c *dig.Container) {
	setProvide(c, server.NewServer)
	setProvide(c, connection.NewConnection)
	provideController(c)
	provideUseCase(c)
	provideRepository(c)
}

func setProvide(c *dig.Container, i any) {
	if err := c.Provide(i); err != nil {
		logger.Error(err.Error())
	}
}