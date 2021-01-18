package models

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func  TestGetLoan (t *testing.T) {
	var (
		expectedSQL = `SELECT
			loan_id, loan_code, loan_amount, loan_status, loan_product
		FROM loans.loans
		WHERE 
			\(loan_id = \$1 OR \$1 <= 0\) AND
			\(loan_code = \$2 OR \$2 = ''\) AND
			is_active = 1::bit`
		columns = []string{"loan_id", "loan_code", "loan_amount", "loan_status", "loan_product"}
	)

	testCases := []struct {
		name      string
		wantError bool
		mockRow   *sqlmock.Rows
	}{
		{
			name: "No row in result",
			wantError: false,
			mockRow: sqlmock.NewRows(columns),
		},
		{
			name: "Failed to scan",
			wantError: true,
			mockRow: sqlmock.NewRows(columns).AddRow(1, "ABS023", 200000.00, LoanStatusUnfinished, nil).RowError(1, errors.New("unsupported Scan")),
		},
	}

	conn, mock := testMockSQL(t)
	defer conn.Close()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectQuery(expectedSQL).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnRows(tc.mockRow)

			loan := &Loan{}
			errx := loan.GetLoan(struct {
				ID   int64
				Code string}{},
			)

			if tc.wantError == true {
				assert.NotNil(t, errx)
			} else {
				assert.Nil(t, errx)

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
			}

		})
	}
}

func TestGetLoanList (t *testing.T) {
	var (
		expectedSQL= `
			SELECT
				loan_id, loan_code, loan_amount, loan_status, loan_product
			FROM loans.loans
			WHERE
				\(loan_code = \$3 OR \$3 = ''\) AND
				\(loan_status = \$4 OR \$4 = '00000000-0000-0000-0000-000000000000'::uuid\) AND
				\(loan_product = \$5 OR \$5 = '00000000-0000-0000-0000-000000000000'::uuid\) AND
				is_active = 1::BIT
			LIMIT \$1 OFFSET \$2
		`
		columns = []string{"loan_id", "loan_code", "loan_amount", "loan_status", "loan_product"}
	)

	testCases := []struct {
		name      string
		wantError bool
		mockRow *sqlmock.Rows
		mockError error
	}{
		{
			name: "Success but no rows in result",
			mockRow: sqlmock.NewRows(columns),
		},
		{
			name: "Success with rows in result",
			mockRow: sqlmock.NewRows(columns).AddRow(1, "KBZ023", 200.00, "", ""),
		},
		{
			name:      "Failed to scan",
			wantError: true,
			mockRow:   sqlmock.NewRows(columns).AddRow(1, "KBZ123", 200.00, "", nil).
				RowError(1, fmt.Errorf("failed to scan")).
				CloseError(fmt.Errorf("should be error")),
		},
		{
			name:      "Failed to execute query",
			wantError: true,
			mockRow:   sqlmock.NewRows(columns),
			mockError: fmt.Errorf("failed to execute sql"),
		},
	}

	conn, mock := testMockSQL(t)
	defer conn.Close()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectQuery(expectedSQL).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnRows(tc.mockRow).WillReturnError(tc.mockError).RowsWillBeClosed()

			_, errx := GetLoanList(struct {Pagination
				Code    string
				Product Product
				Status  LoanStatus}{})

			if tc.wantError == true {
				assert.NotNil(t, errx)
			} else {
				assert.Nil(t, errx)

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
			}

		})
	}

}