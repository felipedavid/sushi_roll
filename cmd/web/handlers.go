package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felipedavid/sushi_roll/internal/models"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (a *app) homePage(w http.ResponseWriter, r *http.Request) {
	games, err := a.game.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "home.tmpl", data)
}

func (a *app) createGame(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	// TODO: Implementar um módulo de validação de formulários e reescrever esta validação
	title := r.Form.Get("title")
	desc := r.Form.Get("description")
	createdAt := "2022-12-12"

	id, err := a.game.Insert(title, desc, createdAt)
	if err != nil {
		a.serverError(w, err)
		return
	}

	url := fmt.Sprintf("/jogo?id=%d", id)
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func (a *app) viewGame(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	game, err := a.game.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w)
		} else {
			a.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", *game)
}

func (a *app) deleteGame(w http.ResponseWriter, r *http.Request) {
	id, err := a.getID(r)
	if err != nil {
		a.notFound(w)
		return
	}

	err = a.game.Delete(id)
	if err != nil {
		// TODO: Test for specific errors and respond based on them
		a.notFound(w)
		return
	}

	fmt.Fprintf(w, "The game with id = %d was deleted\n", id)
}

func (a *app) loginPage(w http.ResponseWriter, r *http.Request) {
	games, err := a.game.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "login.tmpl", data)
}

func (a *app) logupPage(w http.ResponseWriter, r *http.Request) {
	games, err := a.game.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "logup.tmpl", data)
}

func (a *app) jogosPage(w http.ResponseWriter, r *http.Request) {
	games, err := a.game.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "jogos.tmpl", data)
}
