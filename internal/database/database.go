package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Open(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		fmt.Printf("Error opening the db: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("No response from the db: %s\n", err)
	}

	return db, err
}
