package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handler := cors.Handler(Router())

	return http.ListenAndServe(s.listenAddr, handler)
}

func Router() *mux.Router {

	r := mux.NewRouter()
	// Heartbeat endpoint
	r.HandleFunc("/heartbeat", handlers.Heartbeat).Methods(http.MethodGet)

	// Products endpoints
	r.HandleFunc("/products", handlers.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.Create).Methods(http.MethodPost)

	r.HandleFunc("/search", handlers.Search).Methods(http.MethodGet)
	return r
}
