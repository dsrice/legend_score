package ri

import (
	"github.com/labstack/echo/v4"
	"legend_score/infra/database/models"
)

type UserRepository interface {
	GetLoginID(c echo.Context, loginID string) (*models.User, error)
}