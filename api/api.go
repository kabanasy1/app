package api

import (
	"github.com/gorilla/mux"
	"github.com/kabanasy1/internal/handlers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.DefaultHandler).Methods("GET", "POST")
	return r
}
