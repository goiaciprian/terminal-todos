package config

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"terminal-todos/internal"
	"terminal-todos/internal/database"
)

type ConfigStruct struct {
	MAIN_FOLDER_PATH   string
	DATABASE_FILE_PATH string
	ICON_PATH          string
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
		ICON_PATH:          folderPath + "\\icon.png",
	}
}

func FirstTimeSetup() {
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

		resp, err := http.Get(internal.ICON_URL)
		if err != nil {
			fmt.Println("Image request error")
			panic(err)
		}
		defer resp.Body.Close();

		icon, err := os.Create(Instance.ICON_PATH)
		if err != nil {
			fmt.Println("Icon file create error")
			panic(err)
		}

		_, err = io.Copy(icon, resp.Body)
		if err != nil {
			fmt.Println("Copy icon error")
			panic(err)
		}

		defer icon.Close()
		defer f.Close()
	} else {
		fmt.Println("Database file already created")
	}

	db, err := database.Open(Instance.DATABASE_FILE_PATH)
	if err != nil {
		errMessage := fmt.Sprintf("Database open error %s", err)
		panic(errMessage)
	}

	schemaResponse, err := http.Get(internal.SCHEMA_URL)
	if err != nil {
		errMessage := fmt.Sprintf("Schema request file %s", err)
		panic(errMessage)
	}

	defer schemaResponse.Body.Close()

	schemaSql, err := io.ReadAll(schemaResponse.Body)
	if err != nil {
		errMessage := fmt.Sprintf("Schema file open error %s", err)
		panic(errMessage)
	}

	_, err = db.Exec(string(schemaSql))
	if err != nil {
		errMessage := fmt.Sprintf("Schema exec error %s", err)
		panic(errMessage)
	} else {
		fmt.Println("Done installing")
	}

	defer db.Close()
}
