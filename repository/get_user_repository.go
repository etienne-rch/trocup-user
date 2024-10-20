package repository

import (
	"context"
	"errors"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByID(id string) (*models.User, error) {
	var user models.User

	// Conversion de l'ID string en ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	// Recherche dans MongoDB avec le Clerk ID (qui est un string)
	err = config.UserCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
