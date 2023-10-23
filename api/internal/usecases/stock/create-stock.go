package stock

import (
	"errors"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"
)

type CreateStockUsecase struct {
	Stock models.Stock
}

// Execute is a method to execute the usecase
// It call the validations method and if it's ok, it calls the repository to create a new stock
// It returns the created stock, the status code and an error
func (uc *CreateStockUsecase) Execute() (models.Stock, int, error) {
	err := uc.validations()
	if err != nil {
		return models.Stock{}, http.StatusBadRequest, err
	}

	db, err := database.Connect()
	if err != nil {
		return models.Stock{}, http.StatusInternalServerError, err
	}

	repo := repository.Stock{
		MongoDB: db,
	}

	insertedID, err := repo.Create(uc.Stock)
	if err != nil {
		return models.Stock{}, http.StatusInternalServerError, err
	}

	uc.Stock.ID = insertedID

	return uc.Stock, http.StatusOK, nil
}

// validations is a method to validate the usecase fields
func (uc *CreateStockUsecase) validations() error {
	if uc.Stock.Name == "" {
		return errors.New("name is required")
	}
	if uc.Stock.Price == 0 {
		return errors.New("price is required")
	}
	if uc.Stock.Price < 0 {
		return errors.New("price must be positive")
	}

	return nil
}
