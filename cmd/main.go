package main

import (
	"context"

	"github.com/kabanasy1/app/internal/application"
	"github.com/kabanasy1/app/internal/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	cfg := config.GetConfig()
	server := application.NewServer(cfg, ctx)

	err := application.RunServer(server)
	if err != nil {
		log.Fatal(err)
	}
}
