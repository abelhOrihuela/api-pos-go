package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"pos.com/app/handlers"
	"pos.com/app/helpers"
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
	api := r.PathPrefix("/api").Subrouter()
	publicRoutes(api)
	addSignHandler(api)
	return r
}

func publicRoutes(r *mux.Router) {
	s := r.PathPrefix("/public").Subrouter()

	// Heartbeat endpoint
	s.HandleFunc("/heartbeat", handlers.Heartbeat).Methods(http.MethodGet)
	s.HandleFunc("/login", handlers.Login).Methods(http.MethodPost)

}

// In another file
func addSignHandler(r *mux.Router) {

	s := r.PathPrefix("/pos").Subrouter()

	s.Use(helpers.MiddlewareAuth)

	// authe verification
	s.HandleFunc("/me", handlers.Me).Methods(http.MethodGet)

	// Products endpoints
	s.HandleFunc("/orders", handlers.CreateOrder).Methods(http.MethodPost)
	s.HandleFunc("/orders", handlers.GetOrders).Methods(http.MethodGet)

	s.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet)
	s.HandleFunc("/products", handlers.CreateProduct).Methods(http.MethodPost)
	s.HandleFunc("/search", handlers.Search).Methods(http.MethodGet)

	s.HandleFunc("/categories", handlers.CreateCategory).Methods(http.MethodPost)
	s.HandleFunc("/categories", handlers.GetAllCategories).Methods(http.MethodGet)

	// Users endpoints
	s.HandleFunc("/users", handlers.CreateUser).Methods(http.MethodPost)
}
