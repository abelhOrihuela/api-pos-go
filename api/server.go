package api

import (
	"net/http"

	"github.com/gorilla/mux"
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
	http.Handle("/", Router())
	return http.ListenAndServe(s.listenAddr, nil)
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/heartbeat", handlers.Heartbeat).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.GetAllProducts).Methods(http.MethodGet)
	return r
}
