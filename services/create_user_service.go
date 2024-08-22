package services

import (
	"trocup-user/models"
	"trocup-user/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return repository.CreateUser(user)
}
