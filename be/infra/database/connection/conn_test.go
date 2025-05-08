package connection_test

import (
	"legend_score/infra/database/connection"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConnectionSuite struct {
	suite.Suite
}

func (s *ConnectionSuite) TestGetConnection() {
	conn := connection.NewConnection()

	s.NotNil(s.T(), conn.Conn)

	rows, err := conn.Conn.Query("select count(*) from goose_db_version")
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
	// Set GO_ENV to "test" to use test.env
	os.Setenv("GO_ENV", "test")

	// Print the current working directory to help debug
	dir, err := os.Getwd()
	if err != nil {
		t.Logf("Error getting current directory: %v", err)
	} else {
		t.Logf("Current directory: %s", dir)
	}

	// Check if test.env exists in various locations
	_, err = os.Stat("test.env")
	t.Logf("test.env in current dir exists: %v", err == nil)

	_, err = os.Stat("../test.env")
	t.Logf("../test.env exists: %v", err == nil)

	_, err = os.Stat("../../test.env")
	t.Logf("../../test.env exists: %v", err == nil)

	_, err = os.Stat("../../../test.env")
	t.Logf("../../../test.env exists: %v", err == nil)

	suite.Run(t, new(ConnectionSuite))
}