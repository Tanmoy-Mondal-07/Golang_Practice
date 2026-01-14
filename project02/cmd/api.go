package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)    // rate lemiting
	r.Use(middleware.RequestID) // rate lemeting and tressing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working"))
	})

	// http.ListenAndServe(":8000", r)
	return r
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
