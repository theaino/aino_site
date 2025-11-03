package main

import (
	"ainosite/views/pub"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	http.Server
}

func NewServer(addr string) *Server {
	return &Server{http.Server{
		Handler: mux.NewRouter(),
		Addr: addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}}
}

func (s *Server) Route() {
	s.Handler.(*mux.Router).HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		pub.Index().Render(context.Background(), w)
	})
	s.Handler.(*mux.Router).PathPrefix("/res/").Handler(http.FileServer(http.Dir("")))
}

func (s *Server) ListenAndServe() error {
	log.Printf("Listening on %v\n", s.Addr)
	return s.Server.ListenAndServe()
}

