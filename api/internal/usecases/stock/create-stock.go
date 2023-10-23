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
