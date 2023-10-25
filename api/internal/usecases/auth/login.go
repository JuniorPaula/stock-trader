package auth

import (
	"errors"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoginUsecase struct {
	Auth models.Auth
}

// Execute is a method to execute the usecase
// It call the validations method and if it's ok, it calls the repository to create a new auth
// It returns the created auth, the status code and an error
func (uc *LoginUsecase) Execute() (models.Credentials, int, error) {
	err := uc.validations()
	if err != nil {
		return models.Credentials{}, http.StatusBadRequest, err
	}

	db, err := database.Connect()
	if err != nil {
		return models.Credentials{}, http.StatusInternalServerError, err
	}

	userRepo := repository.User{MongoDB: db}
	user, err := userRepo.GetByEmail(uc.Auth.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Credentials{}, http.StatusUnauthorized, errors.New("invalid credentials")
		}
		return models.Credentials{}, http.StatusInternalServerError, err
	}
	validPassword, err := uc.Auth.PasswordMatch(user.Password, uc.Auth.Password)
	if !validPassword || err != nil {
		return models.Credentials{}, http.StatusUnauthorized, errors.New("invalid credentials")
	}
	credentials, err := models.GenerateToken(user.ID.Hex(), user.Name, user.Email, 24, "auth")
	if err != nil {
		return models.Credentials{}, http.StatusBadRequest, err
	}
	// update user token
	user.Token = credentials.Token
	err = userRepo.Update(user)
	if err != nil {
		return models.Credentials{}, http.StatusInternalServerError, err
	}

	return *credentials, http.StatusOK, nil
}

// validations is a method to execute the validations of the usecase
// It returns an error
func (uc *LoginUsecase) validations() error {
	if uc.Auth.Email == "" {
		return errors.New("email is required")
	}
	if uc.Auth.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
