package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jericogantuangco/bookings/pkg/config"
	"github.com/jericogantuangco/bookings/pkg/handlers"
)

func routes(config *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.App.Home)
	mux.Get("/about", handlers.App.About)

	return mux
}
