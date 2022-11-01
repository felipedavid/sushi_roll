package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *app) routes() http.Handler {
	router := httprouter.New()

	// Arquivos estátios
	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", a.homePage)
	router.HandlerFunc(http.MethodGet, "/games", a.gamesPage)
	router.HandlerFunc(http.MethodPost, "/games", a.createGame)
	router.HandlerFunc(http.MethodDelete, "/games", a.deleteGame)
	router.HandlerFunc(http.MethodGet, "/games/:id", a.viewGame)

	// Autenticação de usuários
	router.HandlerFunc(http.MethodGet, "/user/signup", a.userSignUp)
	router.HandlerFunc(http.MethodPost, "/user/signup", a.userSignUpPost)
	router.HandlerFunc(http.MethodGet, "/user/login", a.userLogin)
	router.HandlerFunc(http.MethodPost, "/user/login", a.userLoginPost)
	router.HandlerFunc(http.MethodPost, "/user/logout", a.userLogoutPost)

	return a.logRequest(router)
}
