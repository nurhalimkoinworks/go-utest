package postgres

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nurhalimkoinworks/go-utest/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetLoan (t *testing.T) {
	const (
		expectedSQL = `
			SELECT 
				loan_id, loan_code, proposed_amount, loan_amount, loan_status, loan_product, created_at, created_by
			FROM loans.loans
			WHERE
				\(loan_id = \$1 OR \$1 <= 0\) AND
				\(loan_code = \$2 OR \$2 = ''\)
			LIMIT 1
		`
	)

	var (
		now = time.Now()
		columns = []string{"loan_id", "loan_code", "proposed_amount", "loan_amount", "loan_status", "loan_product", "created_at", "created_by"}
		row1 = []driver.Value{2, "KBZ002", 10000, 15000, models.LoanStatusUnfinished, nil, now, "unit-test"}
	)

	mock, conn := testSQLMock(t)
	defer conn.Close()

	repo := NewLoanRepository(conn)

	type testCase struct {
		name      string
		args      models.GetLoanArgs
		mockRow   *sqlmock.Rows
		mockError error
		wantError bool
	}

	testCases := []testCase{
		{
			name:    "No row in result",
			mockRow: sqlmock.NewRows(columns),
			args: struct {
				ID   int64
				Code string
			}{ID: 0, Code: ""},
		},
		{
			name: "Failed to scan",
			mockRow: sqlmock.NewRows(columns).AddRow(row1...).RowError(1, errors.New("any error")),
			args: struct {
				ID   int64
				Code string
			}{ID: 0, Code: "KBZ002"}, wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectQuery(expectedSQL).WithArgs(tc.args.ID, tc.args.Code).WillReturnRows(tc.mockRow).WillReturnError(tc.mockError)

			_, errx := repo.Get(tc.args)

			if tc.wantError == false {
				assert.Nil(t, errx)
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
			} else {
				assert.NotNil(t, errx)
			}

		})
	}
}