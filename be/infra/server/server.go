package server

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
	"go.uber.org/dig"
	"legend_score/controllers/ci"
	"legend_score/infra/logger"
	"net/http"
	"os"
)

type Server struct {
	echo *echo.Echo
	Auth ci.AuthController
	User ci.UserController
}

type inServer struct {
	dig.In
	Auth ci.AuthController
	User ci.UserController
}

func NewServer(s inServer) *Server {
	return &Server{
		Auth: s.Auth,
		User: s.User,
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (s *Server) Start() {
	s.echo = echo.New()
	s.echo.Validator = &CustomValidator{validator: validator.New()}

	s.echo.Logger.SetLevel(log.DEBUG)
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	s.routing()

	ge := os.Getenv("GO_ENV")
	if ge == "" {
		ge = "dev"
	}

	err := godotenv.Load(fmt.Sprintf("%s.env", ge))
	if err != nil {
		logger.Error(err.Error())
	}

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}

func (s *Server) routing() {
	// Swagger endpoint
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	api := s.echo.Group("/api")
	v := api.Group("/v1")

	v.POST("/login", s.Auth.Login)

	u := v.Group("/user")
	u.POST("", s.User.CreateUser)
	u.GET("", s.User.GetUsers)
	u.GET("/:user_id", s.User.GetUser)
}