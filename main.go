package main

import (
	"log"

	"github.com/joho/godotenv"
	"pos.com/app/api"
)

func main() {
	godotenv.Load(".env")

	setupDatabase()
	server := api.NewServer(":8000")
	log.Fatal(server.Start())
}

func setupDatabase() {
	//db.Connect()

	// if err := db.Database.AutoMigrate(
	// 	&domain.User{},
	// 	&domain.Category{},
	// 	&domain.Product{},
	// 	&domain.Order{},
	// 	&domain.OrderProduct{}); err != nil {
	// 	log.Fatalln(err)
	// }

	/*

		db.Database.Create(&domain.User{
			Username: "admin@hola.com",
			Email:    "admin@hola.com",
			Password: "secret",
			Role:     "admin",
		})

		db.Database.Create(&domain.Category{
			Name:        "default",
			Description: "default",
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
	*/

}
