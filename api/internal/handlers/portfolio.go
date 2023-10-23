package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	ucPortfolio "stocktrader/internal/usecases/portfolio"
)

// BuyPortfolio is a method to buy a new portfolio
func BuyPortfolio(w http.ResponseWriter, r *http.Request) {
	var p models.Portfolio

	err := helpers.ReadJSON(w, r, &p)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := ucPortfolio.BuyPortfolioUsecase{Portfolio: p}
	u, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, u)
}

// SellPortfolio is a method to sell a portfolio
func SellPortfolio(w http.ResponseWriter, r *http.Request) {
	var p models.Portfolio

	err := helpers.ReadJSON(w, r, &p)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := ucPortfolio.SellPortfolioUsecase{Portfolio: p}
	_, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	var response struct {
		Message string `json:"message"`
	}
	response.Message = "Portfolio sold successfully"

	helpers.WriteJSON(w, http.StatusOK, response)
}
