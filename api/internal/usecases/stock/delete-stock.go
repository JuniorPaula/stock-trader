package stock

import (
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteStockUsecase struct {
	ID primitive.ObjectID
}

// Execute is a method to execute the usecase
// and delete a stock
// It returns the status code and an error
func (uc *DeleteStockUsecase) Execute() (int, error) {
	db, err := database.Connect()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	repo := repository.Stock{
		MongoDB: db,
	}

	err = repo.Delete(uc.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}
