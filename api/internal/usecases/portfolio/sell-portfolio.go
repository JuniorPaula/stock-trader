package portfolio

import (
	"errors"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SellPortfolioUsecase struct {
	Portfolio models.Portfolio
}

// Execute is a method to execute the usecase
// It call the validations method and if it's ok, it calls the repository to create a new portfolio
// It returns the created portfolio, the status code and an error
func (uc *SellPortfolioUsecase) Execute() (models.Portfolio, int, error) {
	err := uc.validations()
	if err != nil {
		return models.Portfolio{}, http.StatusBadRequest, err
	}

	db, err := database.Connect()
	if err != nil {
		return models.Portfolio{}, http.StatusInternalServerError, err
	}

	repo := repository.Portfolio{
		MongoDB: db,
	}

	portfolio, err := repo.GetByID(uc.Portfolio.ID)
	if err != nil {
		return models.Portfolio{}, http.StatusInternalServerError, err
	}

	if portfolio.Quantity < uc.Portfolio.Quantity {
		return models.Portfolio{}, http.StatusBadRequest, errors.New("insufficient quantity")
	}
	portfolio.Quantity -= uc.Portfolio.Quantity

	if portfolio.Quantity == 0 {
		err = repo.Delete(portfolio)
		if err != nil {
			return models.Portfolio{}, http.StatusInternalServerError, err
		}
	} else {
		err = repo.Update(portfolio)
		if err != nil {
			return models.Portfolio{}, http.StatusInternalServerError, err
		}
	}

	totalFounds := portfolio.Price * float64(uc.Portfolio.Quantity)

	repoUser := repository.User{MongoDB: db}
	user, err := repoUser.GetByID(uc.Portfolio.UserID)
	if err != nil {
		return models.Portfolio{}, http.StatusInternalServerError, err
	}

	user.Founds += totalFounds
	err = repoUser.Update(user)
	if err != nil {
		return models.Portfolio{}, http.StatusInternalServerError, err
	}

	return portfolio, http.StatusOK, nil
}

// validations is a method to validate the portfolio
// It returns an error
func (uc *SellPortfolioUsecase) validations() error {
	if uc.Portfolio.UserID == primitive.NilObjectID {
		return errors.New("user_id is required")
	}
	if uc.Portfolio.ID == primitive.NilObjectID {
		return errors.New("id is required")
	}
	if uc.Portfolio.Quantity == 0 {
		return errors.New("quantity is required")
	}

	return nil
}
