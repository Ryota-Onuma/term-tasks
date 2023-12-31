package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	_ "embed"

	"github.com/Ryota-Onuma/term-tasks/src/cli"
)

//go:embed db/schema/*.sql
var schemaFiles embed.FS

//go:embed db/seeds/master/*.sql
var masterDataFiles embed.FS

//go:embed db/seeds/local/*.sql
var localDataFiles embed.FS

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dbDirPath := filepath.Join(homeDir, ".term-tasks", "db")
	sqliteFile := filepath.Join(dbDirPath, "db.sqlite3")
	// sqliteFileがなかったら作成する
	if _, err := os.Stat(sqliteFile); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDirPath, 0755); err != nil {
			log.Fatal(err)
		}
		if _, err := os.Create(sqliteFile); err != nil {
			log.Fatal(err)
		}
		fmt.Println("📁 Created " + sqliteFile)
	}
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		log.Fatal(err)
	}

	app := cli.New(db, schemaFiles, masterDataFiles, localDataFiles)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
