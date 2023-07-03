package main

import (
	"github.com/kabanasy1/app/internal/application"
	"github.com/kabanasy1/app/internal/application/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	server := application.NewServer(cfg)

	err := application.RunServer(server)
	if err != nil {
		log.Fatal(err)
	}
}
