package stock

import (
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"
)

type ListStocksUsecase struct{}

func (uc *ListStocksUsecase) Execute() ([]models.Stock, int, error) {
	db, err := database.Connect()
	if err != nil {
		return []models.Stock{}, http.StatusInternalServerError, err
	}

	repo := repository.Stock{MongoDB: db}
	stocks, err := repo.List()
	if err != nil {
		return []models.Stock{}, http.StatusInternalServerError, err
	}

	return stocks, http.StatusOK, nil
}
