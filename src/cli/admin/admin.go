package admin

import (
	"database/sql"
)

type admin struct {
	db *sql.DB
}

func New(db *sql.DB) *admin {
	return &admin{db: db}
}

func (a *admin) DB() *sql.DB {
	return a.db
}
