package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	cfg    config
	logger *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "port on which the server listens")
	flag.StringVar(&cfg.env, "env", "dev", "environment of the server (dev, stg, prod)")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		cfg:    cfg,
		logger: logger,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/healthcheck", app.healthCheck)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	logger.Printf("%s server listening at http://localhost:%d", app.cfg.env, app.cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}

// 586/17
