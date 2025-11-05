package main

import (
	"ainosite/server"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8000"
	}

	path := os.Getenv("SQLITE_PATH")
	if path == "" {
		path = "site.db"
	}

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		b := make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, b); err != nil {
			panic(err)
		}
		sessionKey = base64.RawURLEncoding.EncodeToString(b)
	}
	password := os.Getenv("ADM_PASSWD")
	if password == "" {
		log.Fatal("$ADM_PASSWD is not provided")
	}

	d, err := server.NewDB(path)
	if err != nil {
		log.Fatal(err)
	}
	if err := d.Migrate(); err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(d, addr)
	a := server.NewAuth(sessionKey, password)
	s.Route(a)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
