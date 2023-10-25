package stock

import (
	"math/rand"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/repository"
)

type RandomStockUsecase struct{}

// Execute is a method to execute usecase
// Get the stocks from database, random a stock price and update it
func (uc RandomStockUsecase) Execute() (int, error) {
	db, err := database.Connect()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	repo := repository.Stock{MongoDB: db}
	stocks, err := repo.List()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Random a stock price
	for _, stock := range stocks {
		stock.Price = randomPrice(stock.Price)
		err = repo.Update(stock)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}

// randomPrice is a method to random a price
func randomPrice(price float64) float64 {
	// Random a price between 0.5 and 1.5
	price = price * (0.5 + (1.5-0.5)*rand.Float64())

	return float64(int(price*100)) / 100
}
