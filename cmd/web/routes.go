package main

import "net/http"

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/jogos", a.game)

	return a.logRequest(mux)
}
