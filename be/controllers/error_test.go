package controllers_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"legend_score/consts/ecode"
	"legend_score/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	// Create a new echo instance
	e := echo.New()

	// Test cases
	testCases := []struct {
		name       string
		errorCode  string
		statusCode int
	}{
		{
			name:       "Authentication Error",
			errorCode:  ecode.E0000,
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "Request Error",
			errorCode:  ecode.E0001,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Account Lock Error",
			errorCode:  ecode.E1001,
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "Login ID Already Used Error",
			errorCode:  ecode.E2001,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "System Error",
			errorCode:  ecode.E9000,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Empty Error Code",
			errorCode:  "",
			statusCode: http.StatusInternalServerError, // Should default to E9000
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the ErrorResponse function
			err := controllers.ErrorResponse(c, tc.errorCode)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the status code is correct
			assert.Equal(t, tc.statusCode, rec.Code)

			// Assert that the response body contains the error code
			expectedCode := tc.errorCode
			if expectedCode == "" {
				expectedCode = ecode.E9000
			}
			assert.Contains(t, rec.Body.String(), expectedCode)
		})
	}
}