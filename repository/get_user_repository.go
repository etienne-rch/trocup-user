package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByID(id string) (*models.User, error) {
	var user models.User

	// Recherche dans MongoDB avec le Clerk ID (qui est un string)
	err := config.UserCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
