package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Ryota-Onuma/todo-app/src/cli"
)

func main() {
	sqliteFile, ok := os.LookupEnv("SQLITE_FILE")
	if !ok {
		log.Fatal("SQLITE_FILE is not set")
	}
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		log.Fatal(err)
	}

	app := cli.New(db)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
