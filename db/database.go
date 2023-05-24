package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(databaseName string, isProd bool) {
	var err error

	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	name := databaseName
	passwd := os.Getenv("DB_PASSWD")

	fmt.Printf("ENVIRONMENT: %s \n", os.Getenv("ENV"))

	if isProd {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, passwd, name, port, "require", "America/Mexico_City")
		Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		Database, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

}
