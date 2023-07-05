package api

import (
	"github.com/gorilla/mux"
	"github.com/kabanasy1/app/internal/application/handlers"
)

func Router(u *handlers.UserHandlers, c *handlers.CarHandlers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users/create", u.Create).Methods("POST")
	r.HandleFunc("/api/v1/users/list", u.List).Methods("GET")
	r.HandleFunc("/api/v1/users/find", u.Find).Methods("GET")

	r.HandleFunc("/api/v1/cars/create", c.Create).Methods("POST")
	r.HandleFunc("/api/v1/cars/list", c.List).Methods("GET")
	r.HandleFunc("/api/v1/cars/find", c.Find).Methods("GET")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.WrapNotFound).GetHandler()
	return r
}
