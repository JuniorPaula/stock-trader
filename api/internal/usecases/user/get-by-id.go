package user

import (
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserByIDUsecase struct {
	ID primitive.ObjectID
}

// Execute is a method to get a user by id
func (uc *GetUserByIDUsecase) Execute() (models.User, int, error) {
	db, err := database.Connect()
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	repo := repository.User{MongoDB: db}
	user, err := repo.GetByID(uc.ID)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	portfolioRepo := repository.Portfolio{MongoDB: db}
	portfolios, err := portfolioRepo.GetAllByUserID(uc.ID)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	user.Portfolios = portfolios

	return user, http.StatusOK, nil
}
