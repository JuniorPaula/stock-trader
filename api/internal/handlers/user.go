package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	ucUser "stocktrader/internal/usecases/user"
)

// CreateUser is a method to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := helpers.ReadJSON(w, r, &user)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := ucUser.CreateUserUsecase{User: user}
	u, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, u)
}
