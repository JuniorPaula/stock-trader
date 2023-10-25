package portfolio

import (
	"errors"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"

	ucUser "stocktrader/internal/usecases/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuyPortfolioUsecase struct {
	Portfolio models.Portfolio
}

// Execute is a method to execute the usecase
// It call the validations method and if it's ok, it calls the repository to create a new portfolio
// It returns the created portfolio, the status code and an error
func (uc *BuyPortfolioUsecase) Execute() (models.Portfolio, int, error) {
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

	repoStock := repository.Stock{
		MongoDB: db,
	}
	stock, err := repoStock.GetByID(uc.Portfolio.StockID)
	if err != nil {
		return models.Portfolio{}, http.StatusInternalServerError, err
	}

	totalPrice := stock.Price * float64(uc.Portfolio.Quantity)

	repoUser := repository.User{MongoDB: db}

	userUsercase := ucUser.GetUserByIDUsecase{ID: uc.Portfolio.UserID}
	user, statusCode, err := userUsercase.Execute()
	if err != nil {
		return models.Portfolio{}, statusCode, err
	}

	if user.Founds < totalPrice {
		return models.Portfolio{}, http.StatusBadRequest, errors.New("insufficient founds")
	}

	for _, p := range user.Portfolios {
		if p.StockID == uc.Portfolio.StockID {
			p.Quantity += uc.Portfolio.Quantity
			err = repo.Update(p)
			if err != nil {
				return models.Portfolio{}, http.StatusInternalServerError, err
			}

			user.Founds -= totalPrice
			err = repoUser.Update(user)
			if err != nil {
				return models.Portfolio{}, http.StatusInternalServerError, err
			}

			return p, http.StatusCreated, nil
		}
	}

	user.Founds -= totalPrice
	err = repoUser.Update(user)
	if err != nil {
		return models.Portfolio{}, http.StatusInternalServerError, err
	}

	uc.Portfolio.Name = stock.Name
	uc.Portfolio.Price = stock.Price

	insertedID, err := repo.Create(uc.Portfolio)
	if err != nil {
		return models.Portfolio{}, http.StatusInternalServerError, err
	}

	uc.Portfolio.ID = insertedID

	return uc.Portfolio, http.StatusCreated, nil

}

func (uc *BuyPortfolioUsecase) validations() error {
	if uc.Portfolio.UserID == primitive.NilObjectID {
		return errors.New("user_id is required")
	}
	if uc.Portfolio.StockID == primitive.NilObjectID {
		return errors.New("stock_id is required")
	}
	if uc.Portfolio.Quantity == 0 {
		return errors.New("quantity is required")
	}

	return nil
}
