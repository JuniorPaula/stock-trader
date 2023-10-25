package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	"stocktrader/internal/usecases/auth"
	"stocktrader/internal/usecases/user"
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

func SignUp(w http.ResponseWriter, r *http.Request) {
	var inputData models.User

	err := helpers.ReadJSON(w, r, &inputData)
	if err != nil {
		helpers.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	uc := user.CreateUserUsecase{User: inputData}
	user, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, user)
}
