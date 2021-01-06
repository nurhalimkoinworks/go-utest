package models

import "github.com/koinworks/asgard-heimdal/utils/utarray"

type (
	GetScheduleListArgs struct {
		LoanID int64
	}

	RepaymentSchedule struct {
		ID                int64             `json:"id" db:"schedule_id"`
		VersionID         int64             `json:"version_id" db:"loan_repayment_schedule_version_id"`
		InstallmentNumber int               `json:"installment_number" db:"installment_number"`
		DueDate           string            `json:"due_date" db:"due_date"`
		DayPastDue        int               `json:"dpd" db:"day_past_due"`
		Principal         float64           `json:"principal" db:"principal_collectible"`
		Interest          float64           `json:"interest" db:"interest_collectible"`
		Late              float64           `json:"late" db:"late_collectible"`
		PrincipalPaid     float64           `json:"principal_paid" db:"sum_principal_paid"`
		InterestPaid      float64           `json:"interest_paid" db:"sum_interest_paid"`
		LatePaid          float64           `json:"late_paid" db:"sum_late_paid"`
		Status            InstallmentStatus `json:"status" db:"installment_status"`
		StatusEnglish     string            `json:"status_en"`
		StatusIndonesia   string            `json:"status_id"`
	}
)

func (ox RepaymentSchedule) IsUnpaid () bool {
	return utarray.IsExist(ox.Status, []InstallmentStatus{InstallmentStatusUpComming, InstallmentStatusLate, InstallmentStatusPartiallyPaid})
}

func (ox RepaymentSchedule) IsDue () bool {
	return utarray.IsExist(ox.Status, []InstallmentStatus{InstallmentStatusLate, InstallmentStatusPartiallyPaid})
}