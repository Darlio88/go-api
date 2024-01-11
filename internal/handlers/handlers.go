package handlers

import (
	"github.com/darlio88/go-api-test/internal/middlewares"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Middlewares
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middlewares.Authorization)
		router.Get("/coins", GetCoinsBalance)
	})
}
