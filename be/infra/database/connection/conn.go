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

	err = godotenv.Load(fmt.Sprintf("/go/src/app/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

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

	conn, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	boil.SetDB(conn)
	boil.DebugMode = true

	return conn, nil
}