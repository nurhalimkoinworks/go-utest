package service

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/nurhalimkoinworks/go-utest/models"
)

type ILoanRepository interface {
	Get (args models.GetLoanArgs) (row models.Loan, errx serror.SError)
}

type IRepaymentRepository interface {
	GetScheduleList (args models.GetScheduleListArgs) (rows []models.RepaymentSchedule, errx serror.SError)
}