package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router(h http.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h).Methods("GET", "POST")
	return r
}
