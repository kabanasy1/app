package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kabanasy1/app/internal/application/models"
	log "github.com/sirupsen/logrus"
)

type UserHandlers struct {
	message string
}

func NewUserHandler() *UserHandlers {
	u := new(UserHandlers)
	u.message = "Hello from user handler"
	return u
}

func (u *UserHandlers) Create(w http.ResponseWriter, r *http.Request) {
	log.Infof("Incoming request: client: %s method: %s uri: %s user-agent: %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WrapError(w, err)
		return
	}

	m := map[string]interface{}{
		"result": "OK",
		"data":   user,
	}
	WrapOK(w, m)
}

func (u *UserHandlers) List(w http.ResponseWriter, r *http.Request) {
	log.Infof("Incoming request: client: %s method: %s uri: %s user-agent: %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	m := map[string]interface{}{
		"result": "OK",
		"data":   "TODO: the method List user will be here",
	}
	WrapOK(w, m)
}

func (u *UserHandlers) Find(w http.ResponseWriter, r *http.Request) {
	log.Infof("Incoming request: client: %s method: %s uri: %s user-agent: %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	m := map[string]interface{}{
		"result": "OK",
		"data":   "TODO: the method Find user will be here",
	}
	WrapOK(w, m)
}
