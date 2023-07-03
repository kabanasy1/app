package api

import (
	"github.com/gorilla/mux"
	"github.com/kabanasy1/app/internal/application/handlers"
)

func Router(h *handlers.DefaultHandler, u *handlers.UserHandlers, c *handlers.CarHandlers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Handle).Methods("GET", "POST")
	r.HandleFunc("/api/v1/users", u.DefaultUserHandler).Methods("GET", "POST")
	r.HandleFunc("/api/v1/cars", c.DefaultCarHandler).Methods("GET", "POST")
	return r
}
