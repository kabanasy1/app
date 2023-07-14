package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kabanasy1/app/internal/application/models"
	"github.com/kabanasy1/app/internal/application/services"
	log "github.com/sirupsen/logrus"
)

type UserHandlers struct {
	svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandlers {
	u := new(UserHandlers)
	u.svc = svc
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

	err = u.svc.CreateUser(user)
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

	vars := r.URL.Query()
	list, err := u.svc.ListUsers(strings.Trim(vars.Get("name"), "\""))

	if err != nil {
		WrapError(w, err)
		return
	}

	m := map[string]interface{}{
		"result": "OK",
		"data":   list,
	}
	WrapOK(w, m)
}

func (u *UserHandlers) Find(w http.ResponseWriter, r *http.Request) {
	log.Infof("Incoming request: client: %s method: %s uri: %s user-agent: %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())

	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	user, err := u.svc.FindUser(id)
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
