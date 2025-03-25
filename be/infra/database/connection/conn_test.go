package connection_test

import (
	"legend_score/infra/database/connection"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConnectionSuite struct {
	suite.Suite
}

func (s *ConnectionSuite) TestGetConnection() {
	conn := connection.NewConnection()

	s.NotNil(s.T(), conn.Conn)

	rows, err := conn.Conn.Query("select count(*) from schema_migrations")
	defer func() {
		err = rows.Close()
	}()

	s.Require().NoError(err)

	for rows.Next() {
		var count int
		er := rows.Scan(&count)
		s.Require().NoError(er)
		s.GreaterOrEqual(count, 1)
	}
}

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(ConnectionSuite))
}