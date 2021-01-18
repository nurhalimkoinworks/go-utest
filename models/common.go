package models

type (
	AdminParam struct {
		English,
		Indonesian string
	}

	Pagination struct {
		Page,
		Limit int
	}
)

type IsActive int

const (
	IsActiveFalse IsActive = iota
	IsActiveTrue
)