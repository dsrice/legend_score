package server

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"legend_score/controllers/ci"
	"legend_score/infra/logger"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Server struct {
	echo *echo.Echo
	Auth ci.AuthController
}

type inServer struct {
	dig.In
	Auth ci.AuthController
}

func NewServer(s inServer) *Server {
	return &Server{
		Auth: s.Auth,
	}
}

func (s *Server) Start() {
	s.echo = echo.New()
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3100"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	s.routing()

	err := godotenv.Load(fmt.Sprintf("/go/src/app/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		logger.Error(err.Error())
	}

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}

func (s *Server) routing() {
	api := s.echo.Group("/api")
	v := api.Group("/v1")

	v.GET("/tournament", s.Tournament.GetAll)
}