package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(id primitive.ObjectID, user *models.User) (*models.User, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	_, err := config.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)
}
