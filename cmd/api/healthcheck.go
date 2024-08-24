package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: active")
	fmt.Fprintf(w, "version: %s\n", version)
	fmt.Fprintf(w, "environment: %s\n", app.cfg.env)
}
