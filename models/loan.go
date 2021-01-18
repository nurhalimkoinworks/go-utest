package models

import (
	"database/sql"
	"github.com/koinworks/asgard-heimdal/libs/serror"
)

type (
	Loan struct {
		ID         int64      `json:"id" db:"loan_id"`
		Code       string     `json:"code" db:"loan_code"`
		LoanAmount float64    `json:"loan_amount" db:"loan_amount"`
		Status     LoanStatus `json:"status" db:"loan_status"`
		Product    Product    `json:"product" db:"loan_product"`
	}

	GetLoanArgs struct {
		ID   int64
		Code string
	}

	GetLoanListArgs struct {
		Pagination
		Code    string
		Product Product
		Status  LoanStatus
	}
)

func GetLoanList (args GetLoanListArgs) (data *[]Loan, errx serror.SError) {
	q := `
		SELECT
			loan_id, loan_code, loan_amount, loan_status, loan_product
		FROM loans.loans
		WHERE
			(loan_code = $3 OR $3 = '') AND
			(loan_status = $4 OR $4 = '00000000-0000-0000-0000-000000000000'::uuid) AND
			(loan_product = $5 OR $5 = '00000000-0000-0000-0000-000000000000'::uuid) AND
			is_active = 1::BIT
		LIMIT $1 OFFSET $2
		`

	rws, err := db.Queryx(q, args.Limit, args.Page, args.Code, args.Status, args.Product)
	if err != nil {
		return nil, serror.NewFromErrorc(err, "Failed to execute query")
	} else {
		defer rws.Close()
	}

	var result []Loan
	for rws.Next() {
		var row Loan

		err = rws.StructScan(&row)
		if err != nil {
			return nil, serror.NewFromErrorc(err, "Failed to scan")
		}

		result = append(result, row)
	}

	return &result, nil
}

func (ox *Loan) GetLoan (args GetLoanArgs) serror.SError {
	q := `
		SELECT
			loan_id, loan_code, loan_amount, loan_status, loan_product
		FROM loans.loans
		WHERE 
			(loan_id = $1 OR $1 <= 0) AND
			(loan_code = $2 OR $2 = '') AND
			is_active = 1::bit
	`

	err := db.QueryRowx(q, args.ID, args.Code).StructScan(ox)
	if err != nil && err != sql.ErrNoRows {
		return serror.NewFromErrorc(err, "Failed to scan")
	}

	return nil
}