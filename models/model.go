package models

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func InitModels(dbx *sqlx.DB) {
	db = dbx
}