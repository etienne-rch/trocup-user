package services

import (
	"trocup-user/models"
	"trocup-user/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	return CreateUser(user)
}

func ValidateUserCredentials(email, password string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
