package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"
)

func CreateUser(user *models.User) error {
	_, err := config.UserCollection.InsertOne(context.TODO(), user)
	return err
}
