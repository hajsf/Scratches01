//go:build ignore
// +build ignore

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Define() {
	// os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// global.Database = *db
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS Users (id integer not null primary key, name text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
