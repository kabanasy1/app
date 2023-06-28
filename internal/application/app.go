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

	handler := handlers.NewDefaultHandler()

	server := new(Server)
	server.srv.Addr = ":8080"
	server.srv.Handler = api.Router(handler)

	return server
}

func RunServer(s *Server) error {
	log.Println("Server started")

	err := s.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
