package repositories_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"legend_score/infra/database/connection"
	"legend_score/infra/database/models"
	"legend_score/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserTokenRepository_Insert(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserTokenRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a test user token
	now := time.Now()
	userToken := &models.UserToken{
		UserID:       1,
		Token:        "test-token",
		RefreshToken: "test-refresh-token",
	}

	// Set up the mock to expect any query
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect a SELECT query to populate default values
	rows := sqlmock.NewRows([]string{"id", "user_id"}).
		AddRow(1, 1)
	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(rows)

	// Call the Insert method
	err = repo.Insert(c, userToken)

	// Assert that there was no error
	assert.NoError(t, err)

	// Assert that the CreatedAt and UpdatedAt fields were set
	assert.False(t, userToken.CreatedAt.IsZero())
	assert.False(t, userToken.UpdatedAt.IsZero())
	assert.True(t, userToken.CreatedAt.After(now) || userToken.CreatedAt.Equal(now))
	assert.True(t, userToken.UpdatedAt.After(now) || userToken.UpdatedAt.Equal(now))

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserTokenRepository_Insert_Error(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserTokenRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a test user token
	userToken := &models.UserToken{
		UserID:       1,
		Token:        "test-token",
		RefreshToken: "test-refresh-token",
	}

	// Set up the mock to expect any query and return an error
	mock.ExpectExec("INSERT INTO").WillReturnError(sql.ErrConnDone)

	// Call the Insert method
	err = repo.Insert(c, userToken)

	// Assert that there was an error
	assert.Error(t, err)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}