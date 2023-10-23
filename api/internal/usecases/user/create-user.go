package user

import (
	"errors"
	"net/http"
	"stocktrader/internal/database"
	"stocktrader/internal/models"
	"stocktrader/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUsecase struct {
	User models.User
}

// Execute is a method to execute the usecase
// It call the validations method and if it's ok, it calls the repository to create a new user
// It returns the created user, the status code and an error
func (uc *CreateUserUsecase) Execute() (models.User, int, error) {
	err := uc.validations()
	if err != nil {
		return models.User{}, http.StatusBadRequest, err
	}

	db, err := database.Connect()
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	if err := uc.HashPassword(); err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	repo := repository.User{
		MongoDB: db,
	}

	insertedID, err := repo.Create(uc.User)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	uc.User.ID = insertedID

	return uc.User, http.StatusOK, nil
}

func (uc *CreateUserUsecase) validations() error {
	if uc.User.Name == "" {
		return errors.New("name is required")
	}
	if uc.User.Email == "" {
		return errors.New("email is required")
	}
	if uc.User.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *CreateUserUsecase) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.User.Password = string(hashedPassword)

	return nil
}
