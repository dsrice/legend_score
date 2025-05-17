package repositories_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"legend_score/infra/database/connection"
	"legend_score/infra/database/models"
	"legend_score/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserRepository_Get(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create test conditions
	conditions := []qm.QueryMod{
		models.UserWhere.LoginID.EQ("testuser"),
	}

	// Set up the mock to expect any query
	rows := sqlmock.NewRows([]string{"id", "login_id", "name", "password", "change_pass_flag", "error_count", "error_datetime", "lock_datetime", "created_at", "updated_at", "deleted_flg", "deleted_at"}).
		AddRow(1, "testuser", "Test User", "password", false, 0, nil, nil, time.Now(), time.Now(), false, nil)
	mock.ExpectQuery(".*").WillReturnRows(rows)

	// Call the Get method
	users, err := repo.Get(c, conditions)

	// Assert that there was no error
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "testuser", users[0].LoginID)
	assert.Equal(t, "Test User", users[0].Name)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Get_Error(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create test conditions
	conditions := []qm.QueryMod{
		models.UserWhere.LoginID.EQ("testuser"),
	}

	// Set up the mock to expect any query and return an error
	mock.ExpectQuery(".*").WillReturnError(sql.ErrConnDone)

	// Call the Get method
	users, err := repo.Get(c, conditions)

	// Assert that there was an error
	assert.Error(t, err)
	assert.Nil(t, users)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetLoginID(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the mock to expect any query
	rows := sqlmock.NewRows([]string{"id", "login_id", "name", "password", "change_pass_flag", "error_count", "error_datetime", "lock_datetime", "created_at", "updated_at", "deleted_flg", "deleted_at"}).
		AddRow(1, "testuser", "Test User", "password", false, 0, nil, nil, time.Now(), time.Now(), false, nil)
	mock.ExpectQuery(".*").WillReturnRows(rows)

	// Call the GetLoginID method
	user, err := repo.GetLoginID(c, "testuser")

	// Assert that there was no error
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.LoginID)
	assert.Equal(t, "Test User", user.Name)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetLoginID_NotFound(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the mock to expect any query and return empty rows
	rows := sqlmock.NewRows([]string{"id", "login_id", "name", "password", "change_pass_flag", "error_count", "error_datetime", "lock_datetime", "created_at", "updated_at", "deleted_flg", "deleted_at"})
	mock.ExpectQuery(".*").WillReturnRows(rows)

	// Call the GetLoginID method
	user, err := repo.GetLoginID(c, "nonexistent")

	// Assert that there was an error
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "not found", err.Error())

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetLoginID_Error(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the mock to expect any query and return an error
	mock.ExpectQuery(".*").WillReturnError(sql.ErrConnDone)

	// Call the GetLoginID method
	user, err := repo.GetLoginID(c, "testuser")

	// Assert that there was an error
	assert.Error(t, err)
	assert.Nil(t, user)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Insert(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a test user
	now := time.Now()
	user := &models.User{
		LoginID:  "testuser",
		Name:     "テストユーザー",
		Password: "password",
	}

	// Set up the mock to expect any query
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect a SELECT query to populate default values
	rows := sqlmock.NewRows([]string{"id", "change_pass_flag", "error_count", "deleted_flg"}).
		AddRow(1, false, 0, false)
	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(rows)

	// Call the Insert method
	err = repo.Insert(c, user)

	// Assert that there was no error
	assert.NoError(t, err)

	// Assert that the CreatedAt and UpdatedAt fields were set
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
	assert.True(t, user.CreatedAt.After(now) || user.CreatedAt.Equal(now))
	assert.True(t, user.UpdatedAt.After(now) || user.UpdatedAt.Equal(now))

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Insert_Error(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewUserRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a test user
	user := &models.User{
		LoginID:  "testuser",
		Name:     "Test User",
		Password: "password",
	}

	// Set up the mock to expect any query and return an error
	mock.ExpectExec("INSERT INTO").WillReturnError(sql.ErrConnDone)

	// Call the Insert method
	err = repo.Insert(c, user)

	// Assert that there was an error
	assert.Error(t, err)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}