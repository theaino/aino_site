package main

import (
	"ainosite/views/adm"
	"ainosite/views/pub"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	http.Server
	DB *DB
}

func NewServer(db *DB, addr string) *Server {
	return &Server{
		Server: http.Server{
			Handler: mux.NewRouter(),
			Addr: addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout: 15 * time.Second,
		},
		DB: db,
	}
}

func (s *Server) Route(auth *Auth) {
	r := s.Handler.(*mux.Router).StrictSlash(true)

	r.PathPrefix("/res").Handler(http.FileServer(http.Dir("")))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		pub.Index().Render(context.Background(), w)
	})
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		pub.Login().Render(context.Background(), w)
	}).Methods("GET")
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		correct := r.FormValue("password") == auth.Password
		if !correct {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		auth.Login(w, r)
		http.Redirect(w, r, "/adm", http.StatusFound)
	}).Methods("POST")

	protected := r.PathPrefix("/adm").Subrouter()
	protected.Use(auth.Middleware)

	protected.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		adm.Index().Render(context.Background(), w)
	})
	protected.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.Logout(w, r)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}

func (s *Server) ListenAndServe() error {
	log.Printf("Listening on %v\n", s.Addr)
	return s.Server.ListenAndServe()
}

