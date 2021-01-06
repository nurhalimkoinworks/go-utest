package postgres

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nurhalimkoinworks/go-utest/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetScheduleList (t *testing.T) {
	const (
		expectedSQL = `
			SELECT
				s.schedule_id, s.loan_repayment_schedule_version_id,
				COALESCE\(s.day_past_due, -8\) AS day_past_due, COALESCE\(s.due_date, ''\) AS due_date, s.installment_number, s.installment_status,
				s.principal_collectible, s.interest_collectible, s.late_collectible,
				s.sum_principal_paid, s.sum_interest_paid, s.sum_late_paid
			FROM loans.loan_repayment_schedules s 
			JOIN loans.loan_repayment_schedule_version v ON 
				v.loan_repayment_schedule_version_id = s.loan_repayment_schedule_version_id AND
				v.is_current_version = 1::BIT AND v.is_active = 1::BIT
			WHERE 
				s.is_active = 1::BIT AND
				s.deleted_at IS NULL AND
				\(v.loan_id = \$1 OR \$1 <= 0\)
			ORDER BY s.installment_number
		`
	)

	var (
		columns = []string{"schedule_id", "loan_repayment_schedule_version_id", "day_past_due", "due_date", "installment_number", "installment_status",
			"principal_collectible", "interest_collectible", "late_collectible", "sum_principal_paid", "sum_interest_paid", "sum_late_paid",
		}
		row1 = []driver.Value{1, 2, -8, "", 2, models.InstallmentStatusUpComming, 10000.00, 500.00, 0.00, 0.00, 0.00, 0.00}
		row2 = []driver.Value{2, 2, 0, "2020-02-15", 2, nil, 10000.00, 500.00, 0.00, 0.00, 0.00, 0.00}
	)

	mock, conn := testSQLMock(t)
	defer conn.Close()

	repo := NewRepaymentRepository(conn)

	testCases := []struct {
		name      string
		args      models.GetScheduleListArgs
		wantError bool
		mockRow   *sqlmock.Rows
		mockError error
	}{
		{
			name: "Default value of field",
			args: struct{ LoanID int64 }{LoanID: 0},
			wantError: false,
			mockRow: sqlmock.NewRows(columns).AddRow(row1...),
		},
		{
			name: "Failed to scan",
			args: struct{ LoanID int64 }{LoanID: 0},
			wantError: true,
			mockRow: sqlmock.NewRows(columns).AddRow(row2...).RowError(1, errors.New("any error")),
		},
		{
			name: "Failed to execute query",
			args: struct{ LoanID int64 }{LoanID: 1},
			wantError: true,
			mockRow: sqlmock.NewRows(columns),
			mockError: fmt.Errorf("any error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectQuery(expectedSQL).WithArgs(tc.args.LoanID).WillReturnRows(tc.mockRow).WillReturnError(tc.mockError).RowsWillBeClosed()

			_, errx := repo.GetScheduleList(tc.args)

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