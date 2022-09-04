package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "This is the home page")
	default:
		w.Header().Set("Allowed", http.MethodPost)
		a.clientError(w, http.StatusMethodNotAllowed)
	}
}

func (a *app) game(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			a.notFound(w)
			return
		}

		fmt.Fprintf(w, "Showing game with id = %d", id)
	default:
		w.Header().Set("Allowed", http.MethodGet)
		a.clientError(w, http.StatusMethodNotAllowed)
	}
}
