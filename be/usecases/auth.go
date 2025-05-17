package usecases

import (
	"encoding/base64"
	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/scrypt"
	"legend_score/consts/ecode"
	"legend_score/entities"
	"legend_score/entities/db"
	"legend_score/infra/database/models"
	"legend_score/infra/logger"
	"legend_score/repositories/ri"
	"legend_score/usecases/ui"
	"os"
	"regexp"
	"strconv"
	"time"
)

type authUseCase struct {
	user      ri.UserRepository
	userToken ri.UserTokenRepository
}

func NewAuthUseCase(user ri.UserRepository, userToken ri.UserTokenRepository) ui.AuthUseCase {
	return &authUseCase{
		user:      user,
		userToken: userToken,
	}
}

func (uc *authUseCase) ValidateLogin(c echo.Context, entity *entities.LoginEntity) error {
	logger.Debug("ValidateLogin start")
	if !uc.ValidatePassword(entity.Password) {
		entity.Code = ecode.E0001
		return errors.New("failed to validate password")
	}

	user, err := uc.user.GetLoginID(c, entity.LoginID)
	if err != nil {
		logger.Error(err.Error())
		entity.Code = ecode.E0001
		return err
	}

	nt := time.Now()
	lt := nt.Add(-10 * time.Minute)
	if user.LockDatetime.Valid && user.LockDatetime.Time.After(lt) {
		entity.Code = ecode.E1001
		return errors.New("failed to validate login")
	}

	entity.User = db.UserEntity{}
	entity.User.SetEntity(user)

	logger.Debug("ValidateLogin end")
	return nil
}

func (uc *authUseCase) ValidatePassword(password string) bool {
	hasMinLength := len(password) >= 6

	// 正規表現でそれぞれ条件を確認
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password) // 大文字
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password) // 小文字
	hasNumber := regexp.MustCompile(`\d`).MatchString(password)

	isValid := hasMinLength && hasUpper && hasLower && hasNumber

	return isValid
}

func (uc *authUseCase) Login(c echo.Context, e *entities.LoginEntity) (*string, error) {
	logger.Debug("Login start")
	salt := "legend_score_salt_dev"
	dk, err := scrypt.Key([]byte(e.Password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E9000
		return nil, err
	}

	if base64.StdEncoding.EncodeToString(dk) != e.User.Password {
		logger.Error("Login failed")
		e.Code = ecode.E0001
		e.User.ErrorCount += 1
		return nil, errors.New("Login failed")
	}

	token, err := uc.CreateToken(e.User.ID, 1)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E0001
		return nil, err
	}

	rt, err := uc.CreateToken(e.User.ID, 2)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E0001
		return nil, err
	}

	ut := models.UserToken{
		UserID:       e.User.ID,
		Token:        token,
		RefreshToken: rt,
	}

	err = uc.userToken.Insert(c, &ut)
	if err != nil {
		logger.Error(err.Error())
		e.Code = ecode.E0001
		return nil, err
	}

	logger.Debug("Login end")
	return &token, nil
}

func (uc *authUseCase) CreateToken(id, exec int) (string, error) {
	var d time.Duration
	if exec == 1 {
		d = 10 * time.Minute
	} else {
		d = 24 * time.Hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "issuer",
		Subject:   "subject",
		Audience:  []string{"audience"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        strconv.Itoa(id),
	})

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}