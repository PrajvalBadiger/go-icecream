package api

import (
	"log"
	"net/http"

	"github.com/PrajvalBadiger/go-icecream/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	listen_addr string
	store       storage.Storage
}

// New_server: create new server
func New_server(listen_addr string, store storage.Storage) *Server {
	return &Server{
		listen_addr: listen_addr,
		store:       store,
	}
}

// Run: start server
func (s *Server) Run() error {
	router := mux.NewRouter()

	// create routes
	router.HandleFunc("/api/flavours", make_HTTP_handle_func(s.handle_flavour))
	router.HandleFunc("/api/flavours/{id}", make_HTTP_handle_func(s.handle_flavour_by_id))

	log.Println("Server runnning on port: ", s.listen_addr)
	return http.ListenAndServe(s.listen_addr, router)
}
