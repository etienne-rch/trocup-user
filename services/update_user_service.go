package services

import (
	"trocup-user/models"
	"trocup-user/repository"
)

func UpdateUser(id string, user *models.User) (*models.User, error) {
	return repository.UpdateUser(id, user)
}
