package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"legend_score/consts/ecode"
	"legend_score/controllers"
	"legend_score/infra/logger"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// JWTMiddleware validates JWT tokens in the Authorization header
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return controllers.ErrorResponse(c, ecode.E0000)
		}

		// Check if the Authorization header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return controllers.ErrorResponse(c, ecode.E0000)
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			logger.Error(err.Error())
			return controllers.ErrorResponse(c, ecode.E0000)
		}

		if !token.Valid {
			return controllers.ErrorResponse(c, ecode.E0000)
		}

		// Extract user ID from token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return controllers.ErrorResponse(c, ecode.E0000)
		}

		userID, err := strconv.Atoi(claims["jti"].(string))
		if err != nil {
			return controllers.ErrorResponse(c, ecode.E0000)
		}

		// Set user ID in context for later use
		c.Set("user_id", userID)

		return next(c)
	}
}