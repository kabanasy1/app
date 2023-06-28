package main

import (
	"github.com/kabanasy1/app/application"
	log "github.com/sirupsen/logrus"
)

func main() {
	server := application.NewServer()

	err := application.RunServer(server)
	if err != nil {
		log.Fatal(err)
	}
}
