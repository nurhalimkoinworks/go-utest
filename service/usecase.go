package service

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/nurhalimkoinworks/go-utest/models"
)

type IRepaymentUseCase interface {
	OutstandingRepayment(args models.OutstandingRepaymentArgs) (resp models.OutstandingRepaymentResp, errx serror.SError)
}