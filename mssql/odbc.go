//go:build ignore
// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	db, err := sql.Open("odbc",
		"DSN=CData Access Source")
	if err != nil {
		log.Fatal(err)
	}

	var (
		ordername string
		freight   string
	)

	rows, err := db.Query("SELECT OrderName, Freight FROM Orders WHERE ShipCity = ?", "New York")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ordername, &freight)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ordername, freight)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
