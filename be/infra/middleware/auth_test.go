package middleware_test

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"legend_score/consts/ecode"
	"legend_score/infra/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestJWTMiddleware(t *testing.T) {
	// Create a new echo instance
	e := echo.New()

	// Create a handler function that will be wrapped by the middleware
	handlerCalled := false
	handler := func(c echo.Context) error {
		handlerCalled = true
		return c.String(http.StatusOK, "success")
	}

	// Create the middleware
	middlewareFunc := middleware.JWTMiddleware(handler)

	// Test cases
	t.Run("Valid JWT Token", func(t *testing.T) {
		// Reset the handler called flag
		handlerCalled = false

		// Create a valid JWT token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["jti"] = "123" // User ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		tokenString, err := token.SignedString([]byte("legend_score"))
		assert.NoError(t, err)

		// Create a request with the token in the Authorization header
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the middleware
		err = middlewareFunc(c)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the handler was called
		assert.True(t, handlerCalled)

		// Assert that the user ID was set in the context
		userID := c.Get("user_id")
		assert.Equal(t, 123, userID)

		// Assert that the status code is OK
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Missing Authorization Header", func(t *testing.T) {
		// Reset the handler called flag
		handlerCalled = false

		// Create a request without an Authorization header
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the middleware
		err := middlewareFunc(c)

		// Assert that there was no error (the middleware returns a JSON response)
		assert.NoError(t, err)

		// Assert that the handler was not called
		assert.False(t, handlerCalled)

		// Assert that the status code is Bad Request
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Assert that the response contains the error code
		assert.Contains(t, rec.Body.String(), ecode.E0001)
	})

	t.Run("Invalid Authorization Format", func(t *testing.T) {
		// Reset the handler called flag
		handlerCalled = false

		// Create a request with an invalid Authorization header
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "InvalidFormat")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the middleware
		err := middlewareFunc(c)

		// Assert that there was no error (the middleware returns a JSON response)
		assert.NoError(t, err)

		// Assert that the handler was not called
		assert.False(t, handlerCalled)

		// Assert that the status code is Bad Request
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Assert that the response contains the error code
		assert.Contains(t, rec.Body.String(), ecode.E0001)
	})

	t.Run("Invalid JWT Token", func(t *testing.T) {
		// Reset the handler called flag
		handlerCalled = false

		// Create a request with an invalid token
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer invalidtoken")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the middleware
		err := middlewareFunc(c)

		// Assert that there was no error (the middleware returns a JSON response)
		assert.NoError(t, err)

		// Assert that the handler was not called
		assert.False(t, handlerCalled)

		// Assert that the status code is Unauthorized
		assert.Equal(t, http.StatusUnauthorized, rec.Code)

		// Assert that the response contains the error code
		assert.Contains(t, rec.Body.String(), ecode.E0000)
	})

	t.Run("Expired JWT Token", func(t *testing.T) {
		// Reset the handler called flag
		handlerCalled = false

		// Create an expired JWT token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["jti"] = "123" // User ID
		claims["exp"] = time.Now().Add(-time.Hour).Unix() // Expired
		tokenString, err := token.SignedString([]byte("legend_score"))
		assert.NoError(t, err)

		// Create a request with the expired token
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the middleware
		err = middlewareFunc(c)

		// Assert that there was no error (the middleware returns a JSON response)
		assert.NoError(t, err)

		// Assert that the handler was not called
		assert.False(t, handlerCalled)

		// Assert that the status code is Unauthorized
		assert.Equal(t, http.StatusUnauthorized, rec.Code)

		// Assert that the response contains the error code
		assert.Contains(t, rec.Body.String(), ecode.E0000)
	})

	t.Run("Invalid User ID in JWT Token", func(t *testing.T) {
		// Reset the handler called flag
		handlerCalled = false

		// Create a JWT token with an invalid user ID
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["jti"] = "invalid" // Invalid user ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		tokenString, err := token.SignedString([]byte("legend_score"))
		assert.NoError(t, err)

		// Create a request with the token
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the middleware
		err = middlewareFunc(c)

		// Assert that there was no error (the middleware returns a JSON response)
		assert.NoError(t, err)

		// Assert that the handler was not called
		assert.False(t, handlerCalled)

		// Assert that the status code is Unauthorized
		assert.Equal(t, http.StatusUnauthorized, rec.Code)

		// Assert that the response contains the error code
		assert.Contains(t, rec.Body.String(), ecode.E0000)
	})
}