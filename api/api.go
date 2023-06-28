package api

import (
	"github.com/gorilla/mux"
	"github.com/kabanasy1/app/internal/application/handlers"
)

func Router(h *handlers.DefaultHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Handle).Methods("GET", "POST")
	return r
}
