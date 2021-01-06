package usecase

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/nurhalimkoinworks/go-utest/models"
	"github.com/nurhalimkoinworks/go-utest/service"
)

type repaymentUseCase struct {
	loanRepo      service.ILoanRepository
	repaymentRepo service.IRepaymentRepository
}

func NewRepaymentUseCase (loanRepo service.ILoanRepository, repaymentRepo service.IRepaymentRepository) service.IRepaymentUseCase {
	return repaymentUseCase{
		loanRepo:      loanRepo,
		repaymentRepo: repaymentRepo,
	}
}

func (ox repaymentUseCase) OutstandingRepayment(args models.OutstandingRepaymentArgs) (resp models.OutstandingRepaymentResp, errx serror.SError) {
	loan, errx := ox.loanRepo.Get(models.GetLoanArgs{ID: args.LoanID})
	if errx != nil {
		return resp, errx
	}

	if !loan.IsOnGoing() {
		errx = serror.New("invalid loan status")
		return resp, errx
	}

	schedules, errx := ox.repaymentRepo.GetScheduleList(models.GetScheduleListArgs{LoanID: args.LoanID})
	if errx != nil {
		return resp, errx
	}

	if len (schedules) == 0 {
		errx = serror.New("invalid schedule")
		return resp, errx
	}

	for _, v := range schedules {
		if !v.IsUnpaid() {
			continue
		}
		resp.Principal += v.Principal
		resp.Interest += v.Interest
		resp.Late += v.Late
		resp.PrincipalPaid += v.PrincipalPaid
		resp.InterestPaid += v.InterestPaid
		resp.LatePaid += v.LatePaid

		if v.IsDue() {
			resp.DuePrincipal = v.Principal - v.PrincipalPaid
			resp.DueInterest = v.Interest - v.InterestPaid
			resp.DueLate = v.Late - v.LatePaid
		}

	}

	resp.RemainingPrincipal = resp.Principal - resp.PrincipalPaid
	resp.RemainingInterest = resp.Interest - resp.InterestPaid
	resp.RemainingLate = resp.Late - resp.LatePaid
	resp.Outstanding = resp.RemainingPrincipal + resp.RemainingInterest + resp.RemainingLate

	return resp, errx
}