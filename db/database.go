package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error

	env := os.Getenv("ENV")

	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	passwd := os.Getenv("DB_PASSWD")

	fmt.Printf("ENVIRONMENT: %s \n", os.Getenv("ENV"))

	if env == "production" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, passwd, name, port, "require", "America/Mexico_City")
		Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if env == "test" {
		Database, err = gorm.Open(sqlite.Open(name), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

}
