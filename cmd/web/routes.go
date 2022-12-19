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
	router.Handle(http.MethodDelete, "/game/:id", a.deleteGame)
	router.HandlerFunc(http.MethodPost, "/game", a.createGame)

	router.Handle(http.MethodGet, "/comment/:id", a.viewComment)
	router.Handle(http.MethodDelete, "/comment/:id", a.deleteComment)
	router.HandlerFunc(http.MethodPost, "/comment", a.createComment)

	router.Handle(http.MethodGet, "/category/:id", a.viewCategory)
	router.Handle(http.MethodDelete, "/category/:id", a.deleteCategory)
	router.HandlerFunc(http.MethodPost, "/category", a.createCategory)

	router.HandlerFunc(http.MethodGet, "/login", a.loginPage)
	router.HandlerFunc(http.MethodGet, "/logup", a.logupPage)

	router.HandlerFunc(http.MethodGet, "/games", a.gamesPage)

	return a.logRequest(router)
}
