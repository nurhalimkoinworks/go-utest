package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

func testSQLMock (t *testing.T) (sqlmock.Sqlmock, *sqlx.DB) {
	t.Helper()

	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return mock, sqlx.NewDb(conn, "postgres")
}