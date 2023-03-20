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
	paramsString := r.URL.Path[len("/v1/movies/"):]
	params := strings.Split(paramsString, "/")

	const idIndex = 0

	switch r.Method {
	case http.MethodGet:
		movieID, err := strconv.Atoi(params[idIndex])
		if err != nil || movieID <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid movie ID")
			return
		}

		movie := data.Movie{
			ID:        1,
			Title:     "Finding Nemo",
			Genres:    []string{"kids", "comedy"},
			CreatedAt: time.Now(),
			Version:   1,
			Runtime:   189,
			Year:      2006,
		}

		err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		fmt.Fprintf(w, "Creating a movie")
	default:
		w.Header().Set("Allow", "GET, POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, http.StatusText(http.StatusMethodNotAllowed))
	}
}
