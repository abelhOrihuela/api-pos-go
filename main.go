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
	db.Database.AutoMigrate(&domain.Order{})
	db.Database.AutoMigrate(&domain.OrderProduct{})
	db.Database.AutoMigrate(&domain.User{})
	db.Database.AutoMigrate(&domain.Category{})

	db.Database.Create(&domain.User{
		Username: "admin@hola.com",
		Email:    "admin@hola.com",
		Password: "secret",
		Role:     "admin",
	})

	flagSeed := false

	if flagSeed {
		fake := faker.New()
		p := fake.Person()

		//migrator := db.Database.Migrator()
		//migrator.DropTable(&domain.Product{})

		x := 128.3456
		price := math.Floor(x*100) / 100

		for i := 0; i < 10; i++ {
			db.Database.Create(&domain.Product{
				Barcode:    p.SSN(),
				Price:      price,
				Name:       p.FirstName(),
				CategoryID: 1,
			})

		}
	}

}
