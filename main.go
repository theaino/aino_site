package main

import (
	"log"

	"aino-spring.com/aino_site/config"
	"aino-spring.com/aino_site/database"
	"aino-spring.com/aino_site/server"
)

func main() {
	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Panic(err)
	}
	db, err := database.NewConnetion(conf)
	if err != nil {
		log.Panic(err)
	}
	db.Setup()
	s, err := server.NewServer(db, conf)
	if err != nil {
		log.Panic(err)
	}
	s.SetupManualPages()
	s.SetupApiPages()
	s.Run(conf.Address)
}
