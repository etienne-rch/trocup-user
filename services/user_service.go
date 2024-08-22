package services

import (
	"errors"
	"log"
	"trocup-user/models"
	"trocup-user/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func UpdateUser(id primitive.ObjectID, updates map[string]interface{}) error {
	bsonUpdates := bson.M{}
	for key, value := range updates {
		bsonUpdates[key] = value
	}

	err := repository.UpdateUser(id, bsonUpdates)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("user not found")
		}
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}
