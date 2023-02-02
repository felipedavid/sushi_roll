package main

import (
	"net/http"
)

func (a *app) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"env":     a.config.env,
			"version": version,
		},
	}

	err := a.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
