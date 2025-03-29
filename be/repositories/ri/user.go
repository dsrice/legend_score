package ri

import (
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"legend_score/infra/database/models"
)

type UserRepository interface {
	Get(c echo.Context, condition []qm.QueryMod) ([]*models.User, error)
	GetLoginID(c echo.Context, loginID string) (*models.User, error)
	Insert(c echo.Context, ut *models.User) error
}