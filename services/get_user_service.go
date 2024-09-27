package services

import (
	"trocup-user/models"
	"trocup-user/repository"
)

func GetUserByID(id string) (*models.User, error) {
	return repository.GetUserByID(id)
}
