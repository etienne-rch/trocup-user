package services

import (
	"trocup-user/models"
	"trocup-user/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user *models.User) error {
    return repository.CreateUser(user)
}

func GetUsers() ([]models.User, error) {
    return repository.GetUsers()
}

func GetUserByID(id primitive.ObjectID) (*models.User, error) {
    return repository.GetUserByID(id)
}