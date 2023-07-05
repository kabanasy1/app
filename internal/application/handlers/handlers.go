package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, er := json.Marshal(m)
	if er != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "noshiff")
	w.WriteHeader(httpStatus)
	fmt.Fprintln(w, string(res))
}

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapNotFound(w http.ResponseWriter, r *http.Request) {
	WrapErrorWithStatus(w, errors.New("not found"), http.StatusNotFound)
}

func WrapOK(w http.ResponseWriter, m map[string]interface{}) {
	res, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(res))
}
