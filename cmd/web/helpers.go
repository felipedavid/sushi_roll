package main

import "net/http"

func (a *app) clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (a *app) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}
