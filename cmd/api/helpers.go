package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (a *app) readJSON(r io.Reader, dest any) error {
	err := json.NewDecoder(r).Decode(dest)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formated JSON (at character %d)", syntaxError.Offset)
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type for field (at character %d)", unmarshalTypeError.Offset)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formated EOF")
		default:
			return err
		}
	}
	return nil
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
