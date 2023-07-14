package application

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kabanasy1/app/api"
	"github.com/kabanasy1/app/internal/application/db"
	"github.com/kabanasy1/app/internal/application/handlers"
	"github.com/kabanasy1/app/internal/application/services"
	"github.com/kabanasy1/app/internal/config"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	cfg    config.Config
	ctx    context.Context
	srv    *http.Server
	dbpool *pgxpool.Pool
}

func NewServer(conf config.Config, ctx context.Context) *Server {
	server := new(Server)
	server.cfg = conf
	server.ctx = ctx
	return server
}

func RunServer(s *Server) error {
	log.Println("Starting server ...")
	var err error
	s.dbpool, err = pgxpool.Connect(s.ctx, s.cfg.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}

	userDb := db.NewUserStorage(s.dbpool)
	userSvc := services.NewUserSvc(userDb)

	userHandler := handlers.NewUserHandler(userSvc)
	carHandler := handlers.NewCarHandler()
	router := api.Router(userHandler, carHandler)

	s.srv = &http.Server{
		Addr:    s.cfg.WebAddr + ":" + s.cfg.WebPort,
		Handler: router,
	}

	log.Printf("Server started on %s:%s", s.cfg.WebAddr, s.cfg.WebPort)
	err = s.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
