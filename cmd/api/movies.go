package main

import (
	"fmt"
	"github.com/felipedavid/sushi_roll/internal/data"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) moviesHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	params := path[3:]

	const idIndex = 0

	switch r.Method {
	case http.MethodGet:
		movieID, err := strconv.Atoi(params[idIndex])
		if err != nil || movieID <= 0 {
			app.notFoundResponse(w, r)
			return
		}

		movie := data.Movie{
			ID:        int64(movieID),
			Title:     "Finding Nemo",
			Genres:    []string{"kids", "comedy"},
			CreatedAt: time.Now(),
			Version:   1,
			Runtime:   189,
			Year:      2006,
		}

		err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	case http.MethodPost:
		var input struct {
			Title   string       `json:"title"`
			Year    int32        `json:"year"`
			Runtime data.Runtime `json:"runtime"`
			Genres  []string     `json:"genres"`
		}

		err := app.readJSON(w, r, &input)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		fmt.Fprintf(w, "%+v\n", input)
	default:
		w.Header().Set("Allow", "GET, POST")
		app.methodNotAllowedResponse(w, r)
	}
}
