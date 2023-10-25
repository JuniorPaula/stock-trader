package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	ucStock "stocktrader/internal/usecases/stock"

	"github.com/go-chi/chi"
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

	helpers.WriteJSON(w, http.StatusCreated, s)
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

// UpdateStock is a handler to call usecase and update a stock
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var stock models.Stock
	stock.ID = helpers.StringToPrimitiveObjectID(id)

	err := helpers.ReadJSON(w, r, &stock)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := ucStock.UpdateStockUsecase{Stock: stock}
	statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil)
}

// DeleteStock is a handler to call usecase and delete a stock
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	uc := ucStock.DeleteStockUsecase{ID: helpers.StringToPrimitiveObjectID(id)}
	statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil)
}

// RandomStock is a handler to call usecase and return a random stock
func RandomStock(w http.ResponseWriter, r *http.Request) {
	uc := ucStock.RandomStockUsecase{}
	statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil)
}
