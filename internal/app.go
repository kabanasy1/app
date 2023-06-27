package application

import (
	"net/http"

	"github.com/kabanasy1/app/api"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	srv *http.Server
}

func NewServer() *Server {
	server := &Server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: api.Router(),
		},
	}
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
