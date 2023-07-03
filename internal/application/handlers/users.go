package handlers

import (
	"fmt"
	"net/http"
)

type UserHandlers struct {
	message string
}

func NewUserHandler() *UserHandlers {
	u := new(UserHandlers)
	u.message = "Hello from user handler"
	return u
}

func (u *UserHandlers) DefaultUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, u.message)
}
