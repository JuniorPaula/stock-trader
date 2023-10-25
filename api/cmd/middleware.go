package main

import (
	"context"
	"errors"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"
	"strings"
)

type contextKey string

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the user from the request header
		// validate the user and add it to the request context
		user, err := authenticate(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextKey("user"), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func authenticate(r *http.Request) (*models.User, error) {
	// get the token from the request header
	// validate the token and get the associated user
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("authorization header is required")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("authorization header format must be Bearer {token}")
	}

	token := headerParts[1]
	if len(token) != 26 {
		return nil, errors.New("token is invalid")
	}
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	userRepo := repository.User{MongoDB: db}
	user, err := userRepo.GetByToken(token)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
