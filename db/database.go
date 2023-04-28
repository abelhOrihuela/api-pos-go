package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(dbname string) {
	var err error

	// databaseName := os.Getenv("DB_NAME")
	Database, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

}
