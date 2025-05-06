package repositories_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"legend_score/infra/database/connection"
	"legend_score/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGameRepository_GetByUserID(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewGameRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the mock to expect any query
	gameDate := time.Now()
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "score", "count", "game_date", "created_at", "updated_at", "deleted_flg", "deleted_at"}).
		AddRow(1, 1, "Game 1", 100, 1, gameDate, time.Now(), time.Now(), false, nil).
		AddRow(2, 1, "Game 2", 200, 2, gameDate, time.Now(), time.Now(), false, nil)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// Call the GetByUserID method
	games, err := repo.GetByUserID(c, 1)

	// Assert that there was no error
	assert.NoError(t, err)
	assert.Len(t, games, 2)
	assert.Equal(t, 1, games[0].UserID)
	assert.Equal(t, "Game 1", games[0].Name.String)
	assert.Equal(t, 100, games[0].Score)
	assert.Equal(t, 1, games[1].UserID)
	assert.Equal(t, "Game 2", games[1].Name.String)
	assert.Equal(t, 200, games[1].Score)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_GetByUserID_Error(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewGameRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the mock to expect any query and return an error
	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrConnDone)

	// Call the GetByUserID method
	games, err := repo.GetByUserID(c, 1)

	// Assert that there was an error
	assert.Error(t, err)
	assert.Nil(t, games)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_GetWithDetails(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewGameRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the mock to expect any query
	gameDate := time.Now()
	gameRows := sqlmock.NewRows([]string{"id", "user_id", "name", "score", "count", "game_date", "created_at", "updated_at", "deleted_flg", "deleted_at"}).
		AddRow(1, 1, "Game 1", 100, 1, gameDate, time.Now(), time.Now(), false, nil)
	mock.ExpectQuery("SELECT").WillReturnRows(gameRows)

	// Set up the mock to expect queries for frames and throws
	frameRows := sqlmock.NewRows([]string{"id", "game_id", "frame_number", "score", "created_at", "updated_at", "deleted_flg", "deleted_at"}).
		AddRow(1, 1, 1, 10, time.Now(), time.Now(), false, nil).
		AddRow(2, 1, 2, 20, time.Now(), time.Now(), false, nil)
	mock.ExpectQuery("SELECT").WillReturnRows(frameRows)

	throwRows := sqlmock.NewRows([]string{"id", "game_id", "frame_id", "throw_number", "pins", "created_at", "updated_at", "deleted_flg", "deleted_at"}).
		AddRow(1, 1, 1, 1, 5, time.Now(), time.Now(), false, nil).
		AddRow(2, 1, 1, 2, 5, time.Now(), time.Now(), false, nil).
		AddRow(3, 1, 2, 1, 10, time.Now(), time.Now(), false, nil)
	mock.ExpectQuery("SELECT").WillReturnRows(throwRows)

	// Call the GetWithDetails method
	game, err := repo.GetWithDetails(c, 1)

	// Assert that there was no error
	assert.NoError(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, 1, game.ID)
	assert.Equal(t, 1, game.UserID)
	assert.Equal(t, "Game 1", game.Name.String)
	assert.Equal(t, 100, game.Score)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_GetWithDetails_Error(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Create a mock connection
	conn := &connection.Connection{
		Conn: db,
	}

	// Create the repository with the mock connection
	repo := repositories.NewGameRepository(conn)

	// Create a test echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the mock to expect any query and return an error
	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrConnDone)

	// Call the GetWithDetails method
	game, err := repo.GetWithDetails(c, 1)

	// Assert that there was an error
	assert.Error(t, err)
	assert.Nil(t, game)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}