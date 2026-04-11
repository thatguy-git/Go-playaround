package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute}

	log.Printf("server started on port %s:", app.config.addr)
	return http.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
