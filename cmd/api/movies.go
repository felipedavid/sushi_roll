package main

import (
	"fmt"
	"github.com/felipedavid/sushi_roll/internal/data"
	"github.com/felipedavid/sushi_roll/internal/validator"
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

		v := validator.New()

		v.Check(input.Title != "", "title", "must be provided")
		v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

		v.Check(input.Year != 0, "year", "must be provided")
		v.Check(input.Year <= 1888, "year", "must be greater than 1888")
		v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")

		v.Check(input.Runtime != 0, "runtime", "must be provided")
		v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

		v.Check(input.Genres != nil, "genres", "must be provided")
		v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
		v.Check(len(input.Genres) <= 5, "genres", "must not contains more than 5 genres")
		v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate genres")

		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		fmt.Fprintf(w, "%+v\n", input)
	default:
		w.Header().Set("Allow", "GET, POST")
		app.methodNotAllowedResponse(w, r)
	}
}
