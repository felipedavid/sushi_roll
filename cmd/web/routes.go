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

	router.Handle(http.MethodGet, "/game/:id", a.viewGame)
	router.Handle(http.MethodDelete, "/games/:id", a.deleteGame)
	router.HandlerFunc(http.MethodPost, "/games", a.createGame)

	router.HandlerFunc(http.MethodGet, "/login", a.loginPage)
	router.HandlerFunc(http.MethodGet, "/logup", a.logupPage)
	router.HandlerFunc(http.MethodGet, "/games", a.gamesPage)

	return a.logRequest(router)
}
