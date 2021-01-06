package postgres

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/nurhalimkoinworks/go-utest/models"
)

const sqlGetSchedule = `
		SELECT
			s.schedule_id, s.loan_repayment_schedule_version_id,
			COALESCE(s.day_past_due, -8) AS day_past_due, COALESCE(s.due_date, '') AS due_date, s.installment_number, s.installment_status,
			s.principal_collectible, s.interest_collectible, s.late_collectible,
			s.sum_principal_paid, s.sum_interest_paid, s.sum_late_paid
		FROM loans.loan_repayment_schedules s 
		JOIN loans.loan_repayment_schedule_version v ON 
			v.loan_repayment_schedule_version_id = s.loan_repayment_schedule_version_id AND
			v.is_current_version = 1::BIT AND v.is_active = 1::BIT
		WHERE 
			s.is_active = 1::BIT AND
        s.deleted_at IS NULL AND	
			(v.loan_id = $1 OR $1 <= 0)
		ORDER BY s.installment_number
	`

func (ox repaymentRepository) GetScheduleList(args models.GetScheduleListArgs) (rows []models.RepaymentSchedule, errx serror.SError) {
	rs, err := ox.db.Queryx(sqlGetSchedule, args.LoanID)
	if err != nil {
		errx = serror.NewFromErrorc(err, "Failed to query get schedule")
		return rows, errx
	}
	defer rs.Close()

	for rs.Next() {
		var row models.RepaymentSchedule

		err = rs.StructScan(&row)
		if err != nil {
			errx = serror.NewFromErrorc(err, "Failed to scan struct")
			return rows, errx
		}

		row.StatusIndonesia = row.Status.Translate().Indonesia
		row.StatusEnglish = row.Status.Translate().English

		rows = append(rows, row)
	}

	return rows, errx
}