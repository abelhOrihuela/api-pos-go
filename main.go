package main

import (
	"log"
	"math"

	"github.com/jaswdr/faker"
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

	flagSeed := true

	if flagSeed {
		fake := faker.New()
		p := fake.Person()

		//migrator := db.Database.Migrator()
		//migrator.DropTable(&domain.Product{})

		x := 128.3456
		price := math.Floor(x*100) / 100

		for i := 0; i < 10; i++ {
			db.Database.Create(&domain.Product{
				// Id:      uint(1),
				Barcode: p.SSN(),
				Price:   price,
				Name:    p.FirstName(),
			})

		}
	}

}
