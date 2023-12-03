package db

import (
	"database/sql"
	"embed"
)

type admin struct {
	db              *sql.DB
	schemaFiles     embed.FS
	masterDataFiles embed.FS
	localDataFiles  embed.FS
}

func New(db *sql.DB, schemaFiles, masterDataFiles, localDataFiles embed.FS) *admin {
	return &admin{db: db, schemaFiles: schemaFiles, masterDataFiles: masterDataFiles, localDataFiles: localDataFiles}
}

func (a *admin) DB() *sql.DB {
	return a.db
}
