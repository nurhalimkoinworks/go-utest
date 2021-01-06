package models

// Admin param group Loan Status
type LoanStatus string

const (
	LoanStatusOnGoing      LoanStatus  = "3b6b220b-2019-4279-8343-cbe6167f7574"
	LoanStatusUnfinished   LoanStatus  = "3b6b220b-2019-4279-8343-cbe6167f7577"
)

func (ox LoanStatus) Translate () (res AdminParam) {
	switch ox {
	case LoanStatusOnGoing:
		res.Indonesia = "Sedang Berjalan"
		res.English = "On Going"
	case LoanStatusUnfinished:
		res.Indonesia = "Belum diselesaikan"
		res.English = "Unfinished"
	}

	return res
}

// Admin param group Loan Product
type LoanProduct string

const (
	LoanProductKoinInvoice LoanProduct = "a1fb40e7-e9c5-11e9-97fa-00163e010bca"
	LoanProductKoinbisnis  LoanProduct = "3b6b220b-2019-4279-8343-cbe6167f7547"
)

func (ox LoanProduct) Translate () (res AdminParam) {
	switch ox {
	case LoanProductKoinbisnis:
		res.English = "Koinbisnis"
		res.Indonesia = "Koinbisnis"
	case LoanProductKoinInvoice:
		res.English = "Koininvoice"
		res.Indonesia = "Koininvoice"
	}

	return res
}