package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/nurhalimkoinworks/go-utest/models"
	"github.com/nurhalimkoinworks/go-utest/service"
)

type loanRepository struct {
	db *sqlx.DB
}

func NewLoanRepository (db *sqlx.DB) service.ILoanRepository {
	return loanRepository{db: db}
}

func (ox loanRepository) Get(args models.GetLoanArgs) (row models.Loan, errx serror.SError) {
	q := `
		SELECT 
			loan_id, loan_code, proposed_amount, loan_amount, loan_status, loan_product, created_at, created_by
		FROM loans.loans
		WHERE
			(loan_id = $1 OR $1 <= 0) AND
			(loan_code = $2 OR $2 = '')
		LIMIT 1
	`

	err := ox.db.QueryRowx(q, args.ID, args.Code).StructScan(&row)
	if err != nil && err != sql.ErrNoRows {
		errx = serror.NewFromErrorc(err, "Failed to struct scan")
		return row, errx
	}

	return row, errx
}