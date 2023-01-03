package main

import (
	"fmt"
	"net/http"
)

func (a *app) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "env: %s\n", a.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
