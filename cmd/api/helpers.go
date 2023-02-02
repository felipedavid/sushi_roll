package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (a *app) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (a *app) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	var js []byte
	var err error
	if a.config.env == "prod" {
		js, err = json.Marshal(data)
	} else {
		js, err = json.MarshalIndent(data, "", "\t")
	}
	if err != nil {
		return err
	}

	for key, header := range headers {
		w.Header()[key] = header
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
