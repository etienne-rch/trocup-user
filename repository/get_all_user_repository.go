package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(skip, limit int64) ([]models.User, error) {
	var users []models.User

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := config.UserCollection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
