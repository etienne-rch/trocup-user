package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"
)

// CreateUser inserts a new user into the MongoDB collection with a passed context
func CreateUser(ctx context.Context, user *models.User) error {
	// Set default values for the user
	user.Credit = 0.0
	user.Balance = 0.0
	user.IsPremium = false

	// Use the provided context to insert the user into the MongoDB collection
	_, err := config.UserCollection.InsertOne(ctx, user)
	return err
}
