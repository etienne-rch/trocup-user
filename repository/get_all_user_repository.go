package repository

import (
	"context"
	"time"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetUsers retrieves all users from the database with improved error handling and context management.
func GetUsers() ([]models.User, error) {
	println("GetUsers function started")

	var users []models.User

	// Create a context with a timeout to avoid long-running queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query to find all users
	cursor, err := config.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		println("Error in querying MongoDB:", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	// Check if the cursor contains any documents
	println("Checking cursor")
	if !cursor.Next(ctx) {
		println("No users found in the collection")
	}

	// Use cursor.All to fetch all users at once
	if err = cursor.All(ctx, &users); err != nil {
		println("Error in fetching users:", err.Error())
		return nil, err
	}

	// Log the users array
	println("Users fetched: ", len(users))
	if len(users) == 0 {
		println("No users found.")
	} else {
		println("Users:", users)
	}

	return users, nil
}
