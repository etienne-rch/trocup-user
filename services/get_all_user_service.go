package services

import (
	"trocup-user/models"
	"trocup-user/repository"
)

func GetUsers() ([]models.User, error) {
	return repository.GetUsers()
}
