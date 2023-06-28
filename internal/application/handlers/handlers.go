package handlers

import (
	"fmt"
	"net/http"
)

type DefaultHandler struct {
	message string
}

func NewDefaultHandler() *DefaultHandler {
	def := new(DefaultHandler)
	def.message = "Hello from default handler"
	return def
}

func (def *DefaultHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, def.message)
}
