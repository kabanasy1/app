package application

import (
	"net/http"

	"github.com/kabanasy1/app/api"
	"github.com/kabanasy1/app/internal/application/handlers"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	srv *http.Server
}

func NewServer() *Server {
	server := new(Server)
	return server
}

func RunServer(s *Server) error {
	handler := handlers.NewDefaultHandler()
	router := api.Router(handler)

	s.srv = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started")

	err := s.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
