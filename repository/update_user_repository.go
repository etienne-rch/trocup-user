package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(id string, user *models.User) (*models.User, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	_, err := config.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return GetUserByID(id)  // Utilisation de l'ID string pour la récupération de l'utilisateur mis à jour
}
