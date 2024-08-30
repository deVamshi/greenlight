package main

import (
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	data := envelope{
		"status": "active",
		"system_info": map[string]string{
			"environment": app.cfg.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Server encountered issue while parsing your request", http.StatusInternalServerError)
		return
	}

}
