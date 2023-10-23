package main

import (
	"net/http"
	"stocktrader/internal/handlers"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// users
	mux.Post("/users", handlers.CreateUser)

	// stocks
	mux.Get("/stocks", handlers.ListStocks)
	mux.Post("/stocks", handlers.CreateStock)
	mux.Put("/stocks/{id}", handlers.UpdateStock)
	mux.Delete("/stocks/{id}", handlers.DeleteStock)

	return mux
}
