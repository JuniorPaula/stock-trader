package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	ucStock "stocktrader/internal/usecases/stock"
)

// CreateStock is a handler to call usecase and create a new stock
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := helpers.ReadJSON(w, r, &stock)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := ucStock.CreateStockUsecase{Stock: stock}
	s, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, s)
}

// ListStocks is a handler to call usecase and return a list of stocks
func ListStocks(w http.ResponseWriter, r *http.Request) {
	uc := ucStock.ListStocksUsecase{}
	stocks, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, stocks)
}
