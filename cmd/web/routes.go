package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *app) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", a.homePage)
	router.HandlerFunc(http.MethodGet, "/games", a.viewGame)
	router.HandlerFunc(http.MethodPost, "/games", a.createGame)
	router.HandlerFunc(http.MethodDelete, "/games", a.deleteGame)
	router.HandlerFunc(http.MethodGet, "/login", a.loginPage)
	router.HandlerFunc(http.MethodGet, "/logup", a.logupPage)
	router.HandlerFunc(http.MethodGet, "/jogos", a.jogosPage)

	return a.logRequest(router)
}
