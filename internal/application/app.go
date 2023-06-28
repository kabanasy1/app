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

	server := &Server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: api.Router(h),
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
