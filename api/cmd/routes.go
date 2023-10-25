package main

import (
	"net/http"
	"stocktrader/internal/handlers"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// auth
	mux.Post("/login", handlers.Login)
	mux.Post("/signup", handlers.SignUp)

	mux.Route("/api", func(mux chi.Router) {
		// middleware
		mux.Use(Auth)

		// users
		mux.Get("/users", handlers.ListUsers)
		mux.Post("/users", handlers.CreateUser)
		mux.Get("/users/{id}", handlers.GetUserByID)

		// portfolios
		mux.Post("/buy-portfolio", handlers.BuyPortfolio)
		mux.Post("/sell-portfolio", handlers.SellPortfolio)

		// stocks
		mux.Get("/stocks", handlers.ListStocks)
		mux.Post("/stocks", handlers.CreateStock)
		mux.Put("/stocks/{id}", handlers.UpdateStock)
		mux.Delete("/stocks/{id}", handlers.DeleteStock)
		mux.Get("/stocks/random", handlers.RandomStock)
	})

	return mux
}
