package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//	http://localhost:8080/static/avatar/amyelsner.png
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/api/mysql/customers", app.ListMysqlCustomers)
	mux.Get("/api/mysql/representatives", app.ListMysqlRepresentatives)

	mux.Get("/api/postgres/customers", app.ListPostgresCustomers)
	mux.Get("/api/postgres/representatives", app.ListPostgresRepresentatives)

	return mux
}
