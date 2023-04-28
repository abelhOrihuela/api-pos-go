package main

import (
	"log"

	"github.com/joho/godotenv"
	"pos.com/app/api"
)

func main() {
	godotenv.Load(".env")
	//setupDatabase()
	server := api.NewServer("localhost:3000")
	log.Fatal(server.Start())
}
