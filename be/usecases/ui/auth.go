package ui

import (
	"github.com/labstack/echo/v4"
	"legend_score/entities"
)

type AuthUseCase interface {
	// ValidateLogin
	// ログイン時のバリデーション
	ValidateLogin(c echo.Context, entity *entities.LoginEntity) error

	// ValidatePassword
	// パスワード確認
	ValidatePassword(password string) bool

	// Login
	// 認証処理
	Login(c echo.Context, e *entities.LoginEntity) (*string, error)
}