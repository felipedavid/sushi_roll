package main

import (
    "net/http"
    "fmt"
    "github.com/julienschmidt/httprouter"
    "strconv"
)

func (a *app) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Creating a movie")
}

func (a *app) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    id, err := readIDParam(r)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "show the details of movie %d\n", id)
}
