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
