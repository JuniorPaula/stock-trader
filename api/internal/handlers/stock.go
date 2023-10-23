package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	ucStock "stocktrader/internal/usecases/stock"
)

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := helpers.ReadJSON(w, r, &stock)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	uc := ucStock.CreateStockUsecase{Stock: stock}
	s, statusCode, err := uc.Execute()
	if err != nil {
		helpers.WriteJSON(w, statusCode, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, s)
}
