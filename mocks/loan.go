package mocks

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/nurhalimkoinworks/go-utest/models"
	"github.com/stretchr/testify/mock"
)

type Loan struct {
	mock.Mock
}

func (obj *Loan) GetLoan (args models.GetLoanArgs) serror.SError {
	ret := obj.Called(args)

	var r1 serror.SError
	if rf, ok := ret.Get(1).(func(loanArgs models.GetLoanArgs) serror.SError); ok {
		r1 = rf(args)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(serror.SError)
		}
	}

	return r1
}

func (obj *Loan) GetLoanList(args models.GetLoanListArgs) (*[]models.Loan, serror.SError) {
	var (
		ret = obj.Called(args)
		r0 *[]models.Loan
		r1 serror.SError
	)

	{
		if rf, ok := ret.Get(0).(func(loanArgs models.GetLoanListArgs) *[]models.Loan); ok {
			r0 = rf(args)
		} else {
			if ret.Get(0) != nil {
				r0 = ret.Get(0).(*[]models.Loan)
			}
		}
	}

	{
		if rf, ok := ret.Get(1).(func(loanArgs models.GetLoanListArgs) serror.SError); ok {
			r1 = rf(args)
		} else {
			if ret.Get(1) != nil {
				r1 = ret.Get(1).(serror.SError)
			}
		}
	}

	return r0, r1
}