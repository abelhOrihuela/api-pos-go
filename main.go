package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pos.com/app/handlers"
)

// import (
// 	"log"

// 	"github.com/joho/godotenv"
// 	"pos.com/app/api"
// )

// func main() {
// 	godotenv.Load(".env")

// 	server := api.NewServer("localhost:3000")
// 	log.Fatal(server.Start())
// }

func main() {

	http.Handle("/", Router())

	// mainRouter := mux.NewRouter().StrictSlash(true)
	// mainRouter.HandleFunc("/test/{mystring}", GetRequest).Name("/test/{mystring}").Methods("GET")
	// http.Handle("/", mainRouter)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Something is wrong : " + err.Error())
	}
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/heartbeat", handlers.Heartbeat)
	r.HandleFunc("/products", handlers.GetAllProducts)
	//(...)
	return r
}
