package repositories_test

import (
	"database/sql"
	"errors"
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

func TestGameRepository_GetFrameByGameIDAndFrameCount(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"id", "user_id", "game_id", "frame_count", "frame_score", "strike_flag", "spare_flag", "created_at", "updated_at", "deleted_flg", "deleted_at"}).
		AddRow(1, 1, 1, 1.0, 10, true, false, time.Now(), time.Now(), false, nil)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// Call the GetFrameByGameIDAndFrameCount method
	frame, err := repo.GetFrameByGameIDAndFrameCount(c, 1, 1)

	// Assert that there was no error
	assert.NoError(t, err)
	assert.NotNil(t, frame)
	assert.Equal(t, 1, frame.ID)
	assert.Equal(t, 1, frame.UserID)
	assert.Equal(t, 1, frame.GameID)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_GetFrameByGameIDAndFrameCount_NotFound(t *testing.T) {
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

	// Set up the mock to expect any query and return no rows
	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

	// Call the GetFrameByGameIDAndFrameCount method
	frame, err := repo.GetFrameByGameIDAndFrameCount(c, 1, 1)

	// Assert that there was no error (since no rows is handled)
	assert.Nil(t, err)
	assert.Nil(t, frame)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_GetFrameByGameIDAndFrameCount_Error(t *testing.T) {
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
	mock.ExpectQuery("SELECT").WillReturnError(errors.New("database error"))

	// Call the GetFrameByGameIDAndFrameCount method
	frame, err := repo.GetFrameByGameIDAndFrameCount(c, 1, 1)

	// Assert that there was an error
	assert.Error(t, err)
	assert.Nil(t, frame)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_CreateFrame(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
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

	// Create a frame to insert
	frame := &models.Frame{
		UserID: 1,
		GameID: 1,
	}

	// Set up the mock to expect any query (using regexp)
	mock.ExpectExec(".*").WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the CreateFrame method
	frameID, err := repo.CreateFrame(c, frame)

	// Assert that there was no error
	assert.NoError(t, err)
	assert.Equal(t, 1, frameID)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_CreateFrame_Error(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
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

	// Create a frame to insert
	frame := &models.Frame{
		UserID: 1,
		GameID: 1,
	}

	// Set up the mock to expect any query and return an error
	mock.ExpectExec(".*").WillReturnError(errors.New("database error"))

	// Call the CreateFrame method
	frameID, err := repo.CreateFrame(c, frame)

	// Assert that there was an error
	assert.Error(t, err)
	assert.Equal(t, 0, frameID)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_RegisterThrow(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
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

	// Create a throw to insert
	throw := &models.Throw{
		UserID:     1,
		GameID:     1,
		FrameID:    1,
		ThrowCount: 1,
		ThrowScore: 6,
		StrikeFlag: false,
		SpareFlag:  false,
		Pin1:       1,
		Pin2:       1,
		Pin3:       1,
		Pin4:       0,
		Pin5:       1,
		Pin6:       0,
		Pin7:       1,
		Pin8:       0,
		Pin9:       1,
		Pin10:      0,
	}

	// Set up the mock to expect any query (using regexp)
	mock.ExpectExec(".*").WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the RegisterThrow method
	err = repo.RegisterThrow(c, throw)

	// Assert that there was no error
	assert.NoError(t, err)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGameRepository_RegisterThrow_Error(t *testing.T) {
	// Create a new mock database connection with QueryMatcherOption
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
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

	// Create a throw to insert
	throw := &models.Throw{
		UserID:     1,
		GameID:     1,
		FrameID:    1,
		ThrowCount: 1,
		ThrowScore: 6,
	}

	// Set up the mock to expect any query and return an error
	mock.ExpectExec(".*").WillReturnError(errors.New("database error"))

	// Call the RegisterThrow method
	err = repo.RegisterThrow(c, throw)

	// Assert that there was an error
	assert.Error(t, err)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}