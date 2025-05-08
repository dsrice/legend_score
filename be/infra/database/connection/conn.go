package connection

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"legend_score/infra/logger"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Connection struct {
	Conn *sql.DB
}

func NewConnection() *Connection {
	c, err := getConnection()

	if err != nil {
		panic(err)
	}

	return &Connection{Conn: c}
}

// getConnection
// DB接続
func getConnection() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	ge := os.Getenv("GO_ENV")
	if ge == "" {
		ge = "dev"
	}

	// Try to load the environment file from various locations
	// First, try the current directory
	envLoaded := false

	// Try test.env first if in test mode
	if ge == "test" {
		err = godotenv.Load("test.env")
		if err == nil {
			envLoaded = true
		}
	}

	// If test.env wasn't loaded or GO_ENV is not "test", try {GO_ENV}.env
	if !envLoaded {
		envFile := fmt.Sprintf("%s.env", ge)
		err = godotenv.Load(envFile)
		if err == nil {
			envLoaded = true
		}
	}

	// If still not loaded, try looking in parent directories
	if !envLoaded {
		// Try one level up
		if ge == "test" {
			err = godotenv.Load("../test.env")
			if err == nil {
				envLoaded = true
			}
		}

		if !envLoaded {
			envFile := fmt.Sprintf("../%s.env", ge)
			err = godotenv.Load(envFile)
			if err == nil {
				envLoaded = true
			}
		}

		// Try two levels up
		if !envLoaded {
			if ge == "test" {
				err = godotenv.Load("../../test.env")
				if err == nil {
					envLoaded = true
				}
			}

			if !envLoaded {
				envFile := fmt.Sprintf("../../%s.env", ge)
				err = godotenv.Load(envFile)
				if err == nil {
					envLoaded = true
				}
			}
		}

		// Try three levels up (be directory)
		if !envLoaded {
			if ge == "test" {
				err = godotenv.Load("../../../test.env")
				if err == nil {
					envLoaded = true
				}
			}

			if !envLoaded {
				envFile := fmt.Sprintf("../../../%s.env", ge)
				err = godotenv.Load(envFile)
				if err == nil {
					envLoaded = true
				}
			}
		}
	}

	// If we still couldn't load any environment file, return an error
	if !envLoaded {
		err = fmt.Errorf("could not load any environment file for GO_ENV=%s", ge)
		logger.Error(err.Error())
		return nil, err
	}

	driver := os.Getenv("GOOSE_DRIVER")
	var conn *sql.DB

	if driver == "postgres" {
		// PostgreSQL connection
		connStr := os.Getenv("GOOSE_DBSTRING")
		conn, err = sql.Open("postgres", connStr)
	} else {
		// Default to MySQL connection
		conf := mysql.Config{
			DBName:               os.Getenv("DATABASE_NAME"),
			User:                 os.Getenv("DATABASE_USER"),
			Passwd:               os.Getenv("DATABASE_PASS"),
			Addr:                 os.Getenv("DATABASE_ADDR"),
			Net:                  "tcp",
			Collation:            "utf8mb4_unicode_ci",
			Loc:                  jst,
			ParseTime:            true,
			AllowNativePasswords: true,
		}
		conn, err = sql.Open("mysql", conf.FormatDSN())
	}

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	boil.SetDB(conn)
	boil.DebugMode = true

	return conn, nil
}