package models

type LoanStatus string
const (
	LoanStatusUnfinished        LoanStatus = "3b6b220b-2019-4279-8343-cbe6167f7577"
	LoanStatusUnderReview       LoanStatus = "fe8876f2-1356-11ea-a090-00163e016d4c"
	LoanStatusPendingDisburse   LoanStatus = "3b6b220b-2019-4279-8343-cbe6167f7571"
	LoanStatusOnGoing           LoanStatus = "3b6b220b-2019-4279-8343-cbe6167f7574"
	LoanStatusPaidOff           LoanStatus = "3b6b220b-2019-4279-8343-cbe6167f7575"

)

func (ox LoanStatus) Get () (res AdminParam) {
	switch ox {
	case LoanStatusUnfinished:
		res.English, res.Indonesian = "Unfinished", "Belum Diselesaikan"
	case LoanStatusUnderReview:
		res.English, res.Indonesian = "Under Review", "Sedang Direview"
	case LoanStatusPendingDisburse:
		res.English, res.Indonesian = "Pending Disburse", "Menunggu Pencairan"
	case LoanStatusOnGoing:
		res.English, res.Indonesian = "On Going", "Sedang Berjalan"
	case LoanStatusPaidOff:
		res.English, res.Indonesian = "Paid Off", "Lunas"
	}

	return res
}

type Product string
const (
	ProductKoinBisnis Product = "3b6b220b-2019-4279-8343-cbe6167f7547"
	ProductKoinInvoice Product = "a1fb40e7-e9c5-11e9-97fa-00163e010bca"
)

func (ox Product) Get () (res AdminParam) {
	switch ox {
	case ProductKoinBisnis:
		res.English, res.Indonesian = "Koinbisnis", "Koinbisnis"
	case ProductKoinInvoice:
		res.English, res.Indonesian = "Koininvoice", "Koininvoice"
	}

	return res
}