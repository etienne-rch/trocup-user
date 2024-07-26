package services

import (
	"trocup-user/models"
	"trocup-user/repository"
)

func CreateUser(user *models.User) error {
    return repository.CreateUser(user)
}

func GetUsers() ([]models.User, error) {
    return repository.GetUsers()
}
