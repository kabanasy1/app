package main

import (
	"github.com/kabanasy1/app/internal/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	server := app.NewServer()

	err := app.RunServer(server)
	if err != nil {
		log.Fatal(err)
	}
}
