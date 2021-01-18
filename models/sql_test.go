package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

func testMockSQL (t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection. Test Name: %s. Error: %+v", t.Name(), err)
	}

	dbx := sqlx.NewDb(sqlDB, "sqlmock")
	InitModels(dbx)

	return dbx, mock
}