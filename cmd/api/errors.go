package main

import (
	"fmt"
	"net/http"
)

func (a *app) logError(r *http.Request, err error) {
	a.logger.Println(err)
}

func (a *app) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	err := a.writeJSON(w, status, env, nil)
	if err != nil {
		a.logError(r, err)
		w.WriteHeader(500)
	}
}

func (a *app) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.logError(r, err)
	message := "the server could not process the request"
	a.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (a *app) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the request resource could not be found"
	a.errorResponse(w, r, http.StatusNotFound, message)
}

func (a *app) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the method %s is not supported by this resource", r.Method)
	a.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
