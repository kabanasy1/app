package handlers

import (
	"fmt"
	"net/http"
)

type CarHandlers struct {
	message string
}

func NewCarHandler() *CarHandlers {
	u := new(CarHandlers)
	u.message = "Hello from car handler"
	return u
}

func (u *CarHandlers) DefaultCarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, u.message)
}
