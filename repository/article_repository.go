package repository

import (
	"context"
	"fmt"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
)

func UserExists(userID string) (bool, error) {
	filter := bson.M{"_id": userID}
	count, err := config.UserCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UpdateUserArticleCredit(userID string, articleID string, price float64) (*models.User, error) {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$inc":  bson.M{"credit": price},
		"$push": bson.M{"articles": articleID},
	}

	result, err := config.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("user not found with ID: %s", userID)
	}

	return GetUserByID(userID)
} 