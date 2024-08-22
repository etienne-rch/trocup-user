package services

import (
	"trocup-user/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(id primitive.ObjectID) error {
	return repository.DeleteUser(id)
}
