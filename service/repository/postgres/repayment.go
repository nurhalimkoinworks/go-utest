package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurhalimkoinworks/go-utest/service"
)

type repaymentRepository struct {
	db *sqlx.DB
}

func NewRepaymentRepository (db *sqlx.DB) service.IRepaymentRepository {
	return repaymentRepository{db: db}
}
