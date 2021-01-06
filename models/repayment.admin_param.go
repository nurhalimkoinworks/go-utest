package models

type InstallmentStatus string

const (
	InstallmentStatusUpComming      InstallmentStatus = "95872c70-de79-11e9-97fa-00163e010bca"
	InstallmentStatusLate           InstallmentStatus = "95872c73-de79-11e9-97fa-00163e010bca"
	InstallmentStatusPartiallyPaid  InstallmentStatus = "95872c72-de79-11e9-97fa-00163e010bca"
	InstallmentStatusLateButPaid    InstallmentStatus = "95872c74-de79-11e9-97fa-00163e010bca"
	InstallmentStatusPaid           InstallmentStatus = "95872c71-de79-11e9-97fa-00163e010bca"
)

func (ox InstallmentStatus) Translate () (res AdminParam) {
	switch ox {
	case InstallmentStatusLate:
		res.English = "Late"
		res.Indonesia = "Terlambat"
	case InstallmentStatusLateButPaid:
		res.English = "Late but Paid"
		res.Indonesia = "Terlambat, Dibayar"
	case InstallmentStatusUpComming:
		res.English = "Upcoming"
		res.Indonesia = "Mendatang"
	case InstallmentStatusPartiallyPaid:
		res.English = "Partially Paid"
		res.Indonesia = "Dibayar Sebagian"
	case InstallmentStatusPaid:
		res.English = "Paid"
		res.Indonesia = "Dibayar"

	}

	return res
}

type PaymentType string

const (
	PaymentTypeInstallment 	PaymentType = "3e14a4f0-e4e8-11e9-97fa-00163e010bca"
	PaymentTypeLumpsum 		PaymentType = "3e14a4f1-e4e8-11e9-97fa-00163e010bca"
)

func (ox PaymentType) Translate() (res AdminParam) {
	switch ox {
	case PaymentTypeInstallment:
		res.English = "Installment"
		res.Indonesia = "Cicilan"
	case PaymentTypeLumpsum:
		res.English = "Lumpsum"
		res.Indonesia = "Jumlah Bulat"
	}

	return res
}
