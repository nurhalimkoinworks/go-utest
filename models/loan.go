package models

import "time"

type (
	Loan struct {
		ID             int64       `json:"id" db:"loan_id"`
		Code           string      `json:"code" db:"loan_code"`
		ProposedAmount float64     `json:"proposed_amount" db:"proposed_amount"`
		LoanAmount     float64     `json:"loan_amount" db:"loan_amount"`
		Status         LoanStatus  `json:"status" db:"loan_status"`
		Product        LoanProduct `json:"product" db:"loan_product"`
		CreatedAt      time.Time   `json:"created_at" db:"created_at"`
		CreatedBy      string      `json:"created_by" db:"created_by"`
	}

	GetLoanArgs struct {
		ID   int64
		Code string
	}
)

func (ox Loan) IsOnGoing () bool {
	return ox.Status == LoanStatusOnGoing
}