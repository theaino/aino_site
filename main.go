package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8000"
	}

	path := os.Getenv("SQLITE_PATH")
	if path == "" {
		path = "site.db"
	}

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		log.Fatal("$SESSION_KEY is not provided")
	}
	password := os.Getenv("ADM_PASSWD")
	if sessionKey == "" {
		log.Fatal("$ADM_PASSWD is not provided")
	}

	d, err := NewDB(path)
	if err != nil {
		log.Fatal(err)
	}
	if err := d.Migrate(); err != nil {
		log.Fatal(err)
	}

	s := NewServer(d, addr)
	a := NewAuth(sessionKey, password)
	s.Route(a)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
