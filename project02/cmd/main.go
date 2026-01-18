package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// app.go##########################################################
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)    // rate lemiting
	r.Use(middleware.RequestID) // rate lemeting and tressing
	r.Use(middleware.Logger)// show logs
	r.Use(middleware.Recoverer) // recover from crashes

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working"))
	})

	// http.ListenAndServe(":8000", r)
	return r
}

func(app *application) run(h http.Handler) error {
srv :=&http.Server{
	Addr: app.config.addr,
Handler: h,
WriteTimeout: time.Second*30,
ReadTimeout: time.Second*10,
IdleTimeout: time.Minute,
}

log.Println("server has started at", app.config.addr)

return srv.ListenAndServe()
}

type application struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

//main.go ##########################################################
func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))
	slog.SetDefault(logger)

	h := api.mount()
	err := api.run(h)
	if err != nil {
		slog.Error("server faild to start", err)
		os.Exit(1)
	}
}