package config

import (
	"errors"
	"fmt"
	"os"

	"terminal-todos/internal/database"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)

type ConfigStruct struct {
	MAIN_FOLDER_PATH   string
	DATABASE_FILE_PATH string
}

var Instance *ConfigStruct

func init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	folderPath := configDir + "\\terminal-todos"

	Instance = &ConfigStruct{
		MAIN_FOLDER_PATH:   folderPath,
		DATABASE_FILE_PATH: folderPath + "\\sql.db",
	}
}

func FirstTimeSetup(migrationFolder string) {
	fmt.Println(Instance)
	if _, err := os.Stat(Instance.DATABASE_FILE_PATH); errors.Is(err, os.ErrNotExist) {

		err := os.MkdirAll(Instance.MAIN_FOLDER_PATH, 0750)
		if err != nil {
			fmt.Println("Error creating the app folder")
			panic(err)
		}

		fmt.Println("Created the app folder")

		f, err := os.Create(Instance.DATABASE_FILE_PATH)
		if err != nil {
			fmt.Println("Error creating the db file")
			panic(err)
		}
		fmt.Println("Database file created")

		defer f.Close()
	} else {
		fmt.Println("Database file already created")
	}

	db, err := database.Open(Instance.DATABASE_FILE_PATH)
	if err != nil {
		return
	}

	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		errMessage := fmt.Sprintf("Migration runner error1 %s", err)
		panic(errMessage)
	}

	fileSource, err := (&file.File{}).Open(migrationFolder)
	if err != nil {
		errMessage := fmt.Sprintf("Migration runner error2 %s", err)
		panic(errMessage)
	}

	migrater, err := migrate.NewWithInstance("file", fileSource, "sql", dbDriver)
	if err != nil {
		errMessage := fmt.Sprintf("Migration runner error3 %s", err)
		panic(errMessage)
	}

	if err := migrater.Up(); err != nil {
		errMessage := fmt.Sprintf("Migration runner error4 %s", err)
		panic(errMessage)
	} else {
		fmt.Println("Database migrations done")
	}

	defer db.Close()
	defer dbDriver.Close()
	defer fileSource.Close()
	defer migrater.Close()
}
