package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/felipedavid/sushi_roll/internal/data"
	"github.com/felipedavid/sushi_roll/internal/validator"
)

func (app *application) moviesHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	params := path[3:]

	const idIndex = 0

	switch r.Method {
	case http.MethodGet:
		movieID, err := strconv.ParseInt(params[idIndex], 10, 64)
		if err != nil || movieID <= 0 {
			app.notFoundResponse(w, r)
			return
		}

		movie, err := app.models.Movies.Get(movieID)
		if err != nil {
			if errors.Is(err, data.ErrRecordNotFound) {
				app.notFoundResponse(w, r)
				return
			}
			app.serverErrorResponse(w, r, err)
			return
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

		movie := &data.Movie{
			Title:   input.Title,
			Year:    input.Year,
			Runtime: input.Runtime,
			Genres:  input.Genres,
		}

		v := validator.New()
		data.ValidateMovie(v, movie)

		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		err = app.models.Movies.Insert(movie)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		headers := make(http.Header)
		headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

		err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	case http.MethodPut:
		movieID, err := strconv.ParseInt(params[idIndex], 10, 64)
		if err != nil || movieID <= 0 {
			app.notFoundResponse(w, r)
			return
		}

		movie, err := app.models.Movies.Get(movieID)
		if err != nil {
			if errors.Is(err, data.ErrRecordNotFound) {
				app.notFoundResponse(w, r)
				return
			}
			app.serverErrorResponse(w, r, err)
			return
		}

		var input struct {
			Title   string       `json:"title"`
			Year    int32        `json:"Year"`
			Runtime data.Runtime `json:"Runtime"`
			Genres  []string     `json:"Genres"`
		}

		err = app.readJSON(w, r, &input)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		movie.Title = input.Title
		movie.Year = input.Year
		movie.Runtime = input.Runtime
		movie.Genres = input.Genres

		v := validator.New()

		if data.ValidateMovie(v, movie); !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		err = app.models.Movies.Update(movie)
		if err != nil {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	case http.MethodDelete:
		movieID, err := strconv.ParseInt(params[idIndex], 10, 64)
		if err != nil || movieID <= 0 {
			app.notFoundResponse(w, r)
			return
		}

		err = app.models.Movies.Delete(movieID)
		if err != nil {
			if errors.Is(err, data.ErrRecordNotFound) {
				app.notFoundResponse(w, r)
			} else {
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		err = app.writeJSON(w, http.StatusOK, envelope{"message": "deletion was successful"}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	default:
		w.Header().Set("Allow", "GET, POST")
		app.methodNotAllowedResponse(w, r)
	}
}
