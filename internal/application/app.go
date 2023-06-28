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

	h := handlers.NewDefaultHandler()

	s := new(Server)
	s.srv.Addr = ":8080"
	s.srv.Handler = api.Router(h)

	return s
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
