package main

import (
	"legend_score/di"
	"legend_score/infra/server"

	"go.uber.org/dig"
)

func main() {
	c := dig.New()

	di.BuildContainer(c)
	err := c.Invoke(func(s *server.Server) {
		s.Start()
	})

	if err != nil {
		panic(err)
	}
}