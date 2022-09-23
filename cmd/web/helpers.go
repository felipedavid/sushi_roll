package main

import (
	"net/http"
	"runtime/debug"
)

func (a *app) serverError(w http.ResponseWriter, err error) {
	a.errLog.Printf("%s\n%s\n\n", err.Error(), debug.Stack())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *app) clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (a *app) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}
