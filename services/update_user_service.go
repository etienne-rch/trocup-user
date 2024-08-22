package services

import (
	"trocup-user/models"
	"trocup-user/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(id primitive.ObjectID, user *models.User) (*models.User, error) {
	return repository.UpdateUser(id, user)
}
