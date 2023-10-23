package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	ucPortfolio "stocktrader/internal/usecases/portfolio"
)

// CreatePortfolio is a method to create a new portfolio
func CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	var p models.Portfolio

	err := helpers.ReadJSON(w, r, &p)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := ucPortfolio.CreatePortfolioUsecase{Portfolio: p}
	u, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, u)
}
