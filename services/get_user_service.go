package services

import (
	"trocup-user/models"
	"trocup-user/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	return repository.GetUserByID(id)
}
