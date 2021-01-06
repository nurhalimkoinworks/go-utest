package usecase

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/nurhalimkoinworks/go-utest/models"
	"github.com/nurhalimkoinworks/go-utest/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOutstandingRepayment (t *testing.T) {
	var (
		mockLoanRepo = new(mocks.ILoanRepository)
		mockRepaymentRepo = new(mocks.IRepaymentRepository)
		obj = NewRepaymentUseCase(mockLoanRepo, mockRepaymentRepo)
	)

	testCases := []struct {
		name              string
		mockLoan          models.Loan
		mockLoanError     serror.SError
		mockSchedule      []models.RepaymentSchedule
		mockScheduleError serror.SError
		wantError         bool
	}{
		{
			name: "Should return total remaining loan amount",
			mockLoan: models.Loan{
				Status: models.LoanStatusOnGoing,
			},
			mockSchedule: []models.RepaymentSchedule{
				{
					Principal: 10000.00, Interest: 10.00, Late: 0.00,
					PrincipalPaid: 10000.00, InterestPaid: 10.00, LatePaid: 0.00,
					Status: models.InstallmentStatusPaid, InstallmentNumber: 1, DayPastDue: 0,
				},
				{
					Principal: 10000.00, Interest: 10.00, Late: 5.00,
					PrincipalPaid: 10000.00, InterestPaid: 10.00, LatePaid: 5.00,
					Status: models.InstallmentStatusLateButPaid, InstallmentNumber: 2, DayPastDue: 8,
				},
				{
					Principal: 10000.00, Interest: 10.00, Late: 5.00,
					PrincipalPaid: 5000.00, InterestPaid: 10.00, LatePaid: 5.00,
					Status: models.InstallmentStatusPartiallyPaid, InstallmentNumber: 3, DayPastDue: 9,
				},
				{
					Principal: 10000.00, Interest: 10.00, Late: 1.00,
					PrincipalPaid: 0.00, InterestPaid: 0.00, LatePaid: 0.00,
					Status: models.InstallmentStatusLate, InstallmentNumber: 4, DayPastDue: 3,
				},
				{
					Principal: 10000.00, Interest: 10.00, Late: 0.00,
					PrincipalPaid: 0.00, InterestPaid: 0.00, LatePaid: 0.00,
					Status: models.InstallmentStatusUpComming, InstallmentNumber: 0,
				},
			},
			wantError: false,
		},
		{
			name: "Failed to get schedule list",
			mockLoan: models.Loan{ID: 1, Status: models.LoanStatusOnGoing},
			mockScheduleError: serror.New("any error"),
			wantError: true,
		},
		{
			name: "Should be validated for empty schedule",
			mockLoan: models.Loan{ID: 1, Status: models.LoanStatusOnGoing},
			mockSchedule: []models.RepaymentSchedule{},
			wantError: true,
		},
		{
			name:          "Failed to get loan",
			mockLoanError: serror.New("any error"),
			wantError:     true,
		},
		{
			name: "Should be validated for non-ongoing loan",
			mockLoan: models.Loan{
				Status: models.LoanStatusUnfinished,
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		var (
			expectedResult models.OutstandingRepaymentResp
		)

		if tc.wantError == false {
			for _, v := range tc.mockSchedule {
				if !v.IsUnpaid() {
					continue
				}

				expectedResult.Principal += v.Principal
				expectedResult.Interest += v.Interest
				expectedResult.Late += v.Late
				expectedResult.PrincipalPaid += v.PrincipalPaid
				expectedResult.InterestPaid += v.InterestPaid
				expectedResult.LatePaid += v.LatePaid
				if v.IsDue() {
					expectedResult.DuePrincipal = v.Principal - v.PrincipalPaid
					expectedResult.DueInterest = v.Interest - v.InterestPaid
					expectedResult.DueLate = v.Late - v.LatePaid
				}
			}
			expectedResult.RemainingPrincipal = expectedResult.Principal - expectedResult.PrincipalPaid
			expectedResult.RemainingInterest = expectedResult.Interest - expectedResult.InterestPaid
			expectedResult.RemainingLate = expectedResult.Late - expectedResult.LatePaid

			expectedResult.Outstanding = expectedResult.RemainingPrincipal + expectedResult.RemainingInterest + expectedResult.RemainingLate
		}

		t.Run(tc.name, func(t *testing.T) {
			mockLoanRepo.On("Get", mock.Anything).Return(tc.mockLoan, tc.mockLoanError).Once()
			mockRepaymentRepo.On("GetScheduleList", mock.Anything).Return(tc.mockSchedule, tc.mockScheduleError).Once()

			result, errx := obj.OutstandingRepayment(models.OutstandingRepaymentArgs{})

			if tc.wantError == true {
				assert.NotNil(t, errx)
			} else {
				assert.Equal(t, expectedResult, result)
			}

			//mockLoanRepo.AssertExpectations(t)
			//mockRepaymentRepo.AssertExpectations(t)

		})
	}

}