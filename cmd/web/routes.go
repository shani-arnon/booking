package main

import (
	"net/http"

	"github.com/shaniarnon/booking/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/shaniarnon/booking/pkg/config"
)

func routs(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
