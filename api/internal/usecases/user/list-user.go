package user

import (
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/repository"
)

type ListUserUsecase struct{}

// Execute is a method to list all users
func (uc *ListUserUsecase) Execute() (interface{}, int, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	repo := repository.User{MongoDB: db}
	users, err := repo.List()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return users, http.StatusOK, nil
}
