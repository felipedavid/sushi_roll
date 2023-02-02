package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/felipedavid/sushi_roll/internal/data"
)

func (a *app) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating a movie")
}

func (a *app) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	data := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Stranger Things",
		Year:      2018,
		Runtime:   120,
		Version:   1,
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"movie": data}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
