package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

var userConfig, _ = os.UserConfigDir()
var dir = "\\terminal-todos"
var DB_FILE_PATH = userConfig + dir + "\\sql.db"

func FirstTimeSetup() {
	if _, err := os.Stat(DB_FILE_PATH); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(userConfig+dir, 0750)
		if err != nil {
			fmt.Println("Error creating the app folder")
			panic(err)
		}

		fmt.Println("Created the app folder")

		f, err := os.Create(DB_FILE_PATH)
		if err != nil {
			fmt.Println("Error creating the db file")
			panic(err)
		}
		fmt.Println("Database file created")

		defer f.Close()
	} else {
		fmt.Println("Database file already created")
	}

	db, err := Open()
	if err != nil {
		return
	}

	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		errMessage := fmt.Sprintf("Migration runner error %s", err)
		panic(errMessage)
	}

	fileSource, err := (&file.File{}).Open("file://./../../internal/database/migrations")
	if err != nil {
		errMessage := fmt.Sprintf("Migration runner error %s", err)
		panic(errMessage)
	}

	migrater, err := migrate.NewWithInstance("file", fileSource, "sql", dbDriver)
	if err != nil {
		errMessage := fmt.Sprintf("Migration runner error %s", err)
		panic(errMessage)
	}

	if err := migrater.Up(); err != nil {
		errMessage := fmt.Sprintf("Migration runner error %s", err)
		panic(errMessage)
	} else {
		fmt.Println("Database migrations done")
	}

	defer db.Close()
	defer dbDriver.Close()
	defer fileSource.Close()
	defer migrater.Close()
}


func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DB_FILE_PATH)
	if err != nil {
		fmt.Printf("Error opening the db: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("No response from the db: %s\n", err)
	}

	return db, err
}
