package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8000"
	}

	s := NewServer(addr)
	s.Route()
	panic(s.ListenAndServe())
}
