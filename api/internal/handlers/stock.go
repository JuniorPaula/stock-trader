package handlers

import (
	"log"
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
)

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := helpers.ReadJSON(w, r, &stock)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	log.Println("INFO: create stock:", stock)

	helpers.WriteJSON(w, http.StatusOK, stock)
}
