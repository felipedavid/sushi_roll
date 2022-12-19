package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felipedavid/sushi_roll/internal/models"
	"github.com/julienschmidt/httprouter"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (a *app) homePage(w http.ResponseWriter, r *http.Request) {
	games, err := a.games.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "home.tmpl", data)
}

func (a *app) createGame(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32000)
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	// TODO: Implementar um módulo de validação de formulários e reescrever esta validação
	title := r.Form.Get("title")
	desc := r.Form.Get("description")
	releaseAt := r.Form.Get("release_at")

	_, err = a.games.Insert(title, desc, releaseAt)
	if err != nil {
		a.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "Game created!")
}

func (a *app) viewGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	game, err := a.games.Get(id)
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

func (a *app) deleteGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		a.notFound(w)
		return
	}

	err = a.games.Delete(id)
	if err != nil {
		// TODO: Test for specific errors and respond based on them
		a.notFound(w)
		return
	}

	fmt.Fprintf(w, "The game with id = %d was deleted\n", id)
}

func (a *app) loginPage(w http.ResponseWriter, r *http.Request) {
	games, err := a.games.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "login.tmpl", data)
}

func (a *app) logupPage(w http.ResponseWriter, r *http.Request) {
	games, err := a.games.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "logup.tmpl", data)
}

func (a *app) gamesPage(w http.ResponseWriter, r *http.Request) {
	games, err := a.games.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	data := newTemplateData()
	data.Games = games
	a.render(w, http.StatusOK, "games.tmpl", data)
}

func (a *app) viewComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil || id <= 0 {
		a.serverError(w, err)
		return
	}

	game, err := a.comments.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w)
		} else {
			a.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%#v", game)
}

func (a *app) createComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32000)
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	userID, _ := strconv.ParseInt(r.Form.Get("user_id"), 10, 64)
	gameID, _ := strconv.ParseInt(r.Form.Get("game_id"), 10, 64)
	content := r.Form.Get("content")

	_, err = a.comments.Insert(userID, gameID, content)
	if err != nil {
		a.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "The comment was created!")
}

func (a *app) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		a.notFound(w)
		return
	}

	err = a.comments.Delete(id)
	if err != nil {
		// TODO: Test for specific errors and respond based on them
		a.notFound(w)
		return
	}

	fmt.Fprintf(w, "The comment with id = %d was deleted\n", id)
}

func (a *app) viewCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil || id <= 0 {
		a.serverError(w, err)
		return
	}

	category, err := a.categories.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w)
		} else {
			a.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%#v", category)
}

func (a *app) createCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32000)
	if err != nil {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.Form.Get("title")

	_, err = a.categories.Insert(title)
	if err != nil {
		a.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "The category was created!")
}

func (a *app) deleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		a.notFound(w)
		return
	}

	err = a.categories.Delete(id)
	if err != nil {
		// TODO: Test for specific errors and respond based on them
		a.notFound(w)
		return
	}

	fmt.Fprintf(w, "The category with id = %d was deleted\n", id)
}
