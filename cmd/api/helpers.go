package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err == nil {
		return nil
	}

	var syntaxError *json.SyntaxError
	var unmarshallTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError

	switch {
	case errors.As(err, &syntaxError):
		return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

	case errors.Is(err, io.ErrUnexpectedEOF):
		return fmt.Errorf("body containes badly-formated JSON")

	case errors.As(err, &unmarshallTypeError):
		if unmarshallTypeError.Field != "" {
			return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshallTypeError.Field)
		}
		return fmt.Errorf("body contains incorrect JSON type at character %d", unmarshallTypeError.Offset)

	case errors.Is(err, io.EOF):
		return fmt.Errorf("body must not be empty")

	case errors.As(err, &invalidUnmarshalError):
		panic(err)
	}

	return err
}
