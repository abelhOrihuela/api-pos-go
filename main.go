package main

import (
	"log"

	"github.com/joho/godotenv"
	"pos.com/app/api"
	"pos.com/app/db"
	"pos.com/app/domain"
)

func main() {
	godotenv.Load(".env")
	setupDatabase()
	server := api.NewServer("localhost:3000")
	log.Fatal(server.Start())
}

func setupDatabase() {
	db.Connect()
	db.Database.AutoMigrate(&domain.Product{})
}
