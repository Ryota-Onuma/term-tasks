package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Ryota-Onuma/todo-app/src/cli"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Fail to load .env %v", err)
	}
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
