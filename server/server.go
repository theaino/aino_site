package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Addr string
	Router *chi.Mux
	DB *DB
}

func NewServer(db *DB, addr string) *Server {
	return &Server{
		Addr: addr,
		Router: chi.NewRouter(),
		DB: db,
	}
}


func (s *Server) ListenAndServe() error {
	log.Printf("Listening on %v...", s.Addr)
	return http.ListenAndServe(s.Addr, s.Router)
}

