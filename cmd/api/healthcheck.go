package main

import (
	"net/http"
)

func (a *app) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": a.config.env,
		"version":     version,
	}

	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.logger.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
