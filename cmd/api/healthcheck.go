package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	js := `{"status":"active", "version":%q, "environment":%q}`
	js = fmt.Sprintf(js, version, app.cfg.env)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(js))
}
