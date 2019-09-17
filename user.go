package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func openUserFile() {
	db, err := sql.Open("sqlite3", "./data/shckwave.db3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func ValidUserName(handle string) bool {
	return true
}
