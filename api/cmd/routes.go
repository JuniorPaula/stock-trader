package main

import (
	"net/http"
	"stocktrader/internal/handlers"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/stocks", handlers.ListStocks)
	mux.Post("/stocks", handlers.CreateStock)

	return mux
}
