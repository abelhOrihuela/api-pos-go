package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"pos.com/app/db"
	"pos.com/app/domain"
	"pos.com/app/dto"
	"pos.com/app/handlers"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {

	router := mux.NewRouter()

	db.Connect()
	//db.Database.Create(&domain.Product{Name: "D42", Barcode: "100"})

	router.HandleFunc("/", handlers.GetAllProducts).Methods(http.MethodPost)

	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUser(rw http.ResponseWriter, r *http.Request) {

	// Read
	//var user domain.User

	var response []dto.ProductResponse

	c := make([]domain.Product, 0)

	db.Database.Find(&c) // find product with integer primary key

	for _, cs := range c {
		response = append(response, cs.ToDto())
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}

}
