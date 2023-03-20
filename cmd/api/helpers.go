package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	var js []byte
	var err error
	if app.config.env == "development" {
		js, err = json.MarshalIndent(data, "", "\t")
	} else {
		js, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	for key, val := range headers {
		w.Header()[key] = val
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}
