package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindUserByEmail checks if a user with the given email exists
func FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	
	var user models.User

	// Query MongoDB to find a user with the given email
	err := config.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByPseudo checks if a user with the given pseudo (username) exists
func FindUserByPseudo(ctx context.Context, pseudo string) (*models.User, error) {
	
	var user models.User

	// Query MongoDB to find a user with the given pseudo
	err := config.UserCollection.FindOne(ctx, bson.M{"pseudo": pseudo}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

