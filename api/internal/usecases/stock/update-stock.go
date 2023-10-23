package stock

import (
	"errors"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"
)

type UpdateStockUsecase struct {
	Stock models.Stock
}

// Execute is a method to execute the usecase
// It call the validations method and if it's ok, it calls the repository to update a stock
// It returns the status code and an error
func (uc *UpdateStockUsecase) Execute() (int, error) {
	err := uc.validations()
	if err != nil {
		return http.StatusBadRequest, err
	}

	db, err := database.Connect()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	repo := repository.Stock{
		MongoDB: db,
	}

	err = repo.Update(uc.Stock)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// validations is a method to validate the usecase fields
func (uc *UpdateStockUsecase) validations() error {
	if uc.Stock.ID.IsZero() {
		return errors.New("id is required")
	}
	if uc.Stock.Name == "" {
		return errors.New("name is required")
	}
	if uc.Stock.Price == 0 {
		return errors.New("price is required")
	}

	return nil
}
