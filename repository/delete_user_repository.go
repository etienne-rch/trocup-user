package repository

import (
	"context"
	"trocup-user/config"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(id string) error {
	// Suppression par Clerk ID qui est une chaîne
	_, err := config.UserCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
