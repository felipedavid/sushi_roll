package main

import (
    "fmt"
    "net/http"
)

func (a *app) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "status: available\n")
    fmt.Fprintf(w, "environment: %s\n", a.config.env)
    fmt.Fprintf(w, "version: %s\n", version)
}
