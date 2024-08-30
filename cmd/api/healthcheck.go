package main

import (
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "active",
		"environment": app.cfg.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Server encountered issue while parsing your request", http.StatusInternalServerError)
		return
	}

}
