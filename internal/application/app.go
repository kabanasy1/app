package application

import (
	"net/http"

	"github.com/kabanasy1/app/api"
	"github.com/kabanasy1/app/internal/application/config"
	"github.com/kabanasy1/app/internal/application/handlers"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	cfg config.Config
	srv *http.Server
}

func NewServer(conf config.Config) *Server {
	server := new(Server)
	server.cfg = conf
	return server
}

func RunServer(s *Server) error {
	handler := handlers.NewDefaultHandler()
	userHandler := handlers.NewUserHandler()
	carHandler := handlers.NewCarHandler()
	router := api.Router(handler, userHandler, carHandler)

	s.srv = &http.Server{
		Addr:    s.cfg.WebAddr + ":" + s.cfg.WebPort,
		Handler: router,
	}

	log.Printf("Server started on %s:%s", s.cfg.WebAddr, s.cfg.WebPort)
	err := s.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
