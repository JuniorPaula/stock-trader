package handlers

import (
	"net/http"
	"stocktrader/internal/helpers"
	"stocktrader/internal/models"
	ucUser "stocktrader/internal/usecases/user"

	"github.com/go-chi/chi"
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

// ListUsers is a method to list all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	uc := ucUser.ListUserUsecase{}
	u, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, u)
}

// GetUserByID is a method to get a user by id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userID := helpers.StringToPrimitiveObjectID(id)

	uc := ucUser.GetUserByIDUsecase{ID: userID}
	user, statusCode, err := uc.Execute()
	if err != nil {
		helpers.ErrorJSON(w, statusCode, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, user)
}
