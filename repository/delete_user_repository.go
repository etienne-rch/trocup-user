package repository

import (
	"context"
	"trocup-user/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(id primitive.ObjectID) error {
	_, err := config.UserCollection.DeleteOne(context.TODO(), primitive.M{"_id": id})
	return err
}
