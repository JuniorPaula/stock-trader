package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	"stocktrader/internal/usecases/auth"
)

// Login is a method to handle the login route
// It receives a http.ResponseWriter and a http.Request
// It returns a json response with the credentials and the status code
func Login(w http.ResponseWriter, r *http.Request) {
	var inputAuth models.Auth

	err := helpers.ReadJSON(w, r, &inputAuth)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := auth.LoginUsecase{Auth: inputAuth}
	credentials, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, credentials)
}
