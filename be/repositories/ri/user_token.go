package ri

import (
	"github.com/labstack/echo/v4"
	"legend_score/infra/database/models"
)

type UserTokenRepository interface {
	Insert(c echo.Context, ut *models.UserToken) error
}