package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kabanasy1/app/internal/application/models"
	log "github.com/sirupsen/logrus"
)

type CarHandlers struct {
	message string
}

func NewCarHandler() *CarHandlers {
	c := new(CarHandlers)
	c.message = "Hello from car handler"
	return c
}

func (c *CarHandlers) Create(w http.ResponseWriter, r *http.Request) {
	log.Infof("Incoming request: client: %s method: %s uri: %s user-agent: %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	var car models.Car

	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		WrapError(w, err)
		return
	}

	m := map[string]interface{}{
		"result": "OK",
		"data":   car,
	}
	WrapOK(w, m)
}

func (c *CarHandlers) List(w http.ResponseWriter, r *http.Request) {
	log.Infof("Incoming request: client: %s method: %s uri: %s user-agent: %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	m := map[string]interface{}{
		"result": "OK",
		"data":   "TODO: the method List car will be here",
	}
	WrapOK(w, m)
}

func (c *CarHandlers) Find(w http.ResponseWriter, r *http.Request) {
	log.Infof("Incoming request: client: %s method: %s uri: %s user-agent: %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	m := map[string]interface{}{
		"result": "OK",
		"data":   "TODO: the method Find car will be here",
	}
	WrapOK(w, m)
}
