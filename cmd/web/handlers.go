package main

import (
	"errors"
	"fmt"
	"github.com/felipedavid/sushi_roll/internal/models"
	"net/http"
	"strconv"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		games, err := a.games.Latest()
		if err != nil {
			a.serverError(w, err)
			return
		}

		if len(games) > 0 {
			for _, game := range games {
				fmt.Fprintf(w, "%v\n", *game)
			}
		} else {
			fmt.Fprintf(w, "Nenhum joguinho foi registrado")
		}

	default:
		w.Header().Set("Allowed", http.MethodPost)
		a.clientError(w, http.StatusMethodNotAllowed)
	}
}

func (a *app) game(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
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

		fmt.Fprintf(w, "%v", *game)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			a.clientError(w, http.StatusBadRequest)
			return
		}

		// TODO: Implementar um módulo de validação de formulários e reescrever esta validação
		title := r.Form.Get("title")
		desc := r.Form.Get("description")
		createdAt := "2022-12-12"

		id, err := a.games.Insert(title, desc, createdAt)
		if err != nil {
			a.serverError(w, err)
			return
		}

		url := fmt.Sprintf("/jogo?id=%d", id)
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	default:
		w.Header().Set("Allowed", http.MethodGet)
		a.clientError(w, http.StatusMethodNotAllowed)
	}
}
