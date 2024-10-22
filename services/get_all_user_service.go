package services

import (
	"trocup-user/models"
	"trocup-user/repository"
)

func GetUsers(skip, limit int64) ([]models.User, error) {
	return repository.GetUsers(skip, limit)
}
