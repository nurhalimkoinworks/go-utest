package models

type OutstandingType int

const (
	OutstadingTypeRepayment OutstandingType = iota
	OutstadingTypeEarlyRepayment
)

type (
	OutstandingRepaymentResp struct {
		Principal                 float64 `json:"principal"`
		Interest                  float64 `json:"interest"`
		Late                      float64 `json:"late"`
		PrincipalPaid             float64 `json:"principal_paid"`
		InterestPaid              float64 `json:"interest_paid"`
		LatePaid                  float64 `json:"late_paid"`
		DuePrincipal              float64 `json:"due_principal"`
		DueInterest               float64 `json:"due_interest"`
		DueLate                   float64 `json:"due_late"`
		RemainingPrincipal        float64 `json:"remaining_principal"`
		RemainingInterest         float64 `json:"remaining_interest"`
		RemainingLate             float64 `json:"remaining_late"`
		Outstanding               float64 `json:"outstanding"`
	}

	OutstandingRepaymentArgs struct {
		LoanID           int64
		IsEarlyRepayment OutstandingType
	}
)
