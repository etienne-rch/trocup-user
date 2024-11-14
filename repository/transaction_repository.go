package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUsersTransaction(userA, userB, articleA, articleB string, articlePriceA, articlePriceB float64, isOneToOne bool) (*models.User, error) {
	filterA := bson.M{"_id": userA}
	filterB := bson.M{"_id": userB}

	if isOneToOne {
		// Case 1 to 1: only modify credits and articles
		updateA := bson.M{
			"$inc":  bson.M{"credit": -articlePriceA},  // UserA's credit decreases
			"$pull": bson.M{"articles": articleA}, // Remove UserA's article
		}
		
		updateB := bson.M{
			"$inc":  bson.M{"credit": -articlePriceB},  // UserB's credit decreases
			"$pull": bson.M{"articles": articleB}, // Remove UserB's article
		}

		_, err := config.UserCollection.UpdateOne(context.TODO(), filterA, updateA)
		if err != nil {
			return nil, err
		}

		_, err = config.UserCollection.UpdateOne(context.TODO(), filterB, updateB)
		if err != nil {
			return nil, err
		}
	} else {
		// Case: userA took-1tM
		updateA := bson.M{
			"$inc": bson.M{"balance": -articlePriceB},  // UserA's balance decreases
		}
		
		updateB := bson.M{
			"$inc": bson.M{
				"balance": articlePriceB,   // UserB's balance increases
				"credit": -articlePriceB,   // UserB's credit decreases
			},
			"$pull": bson.M{"articles": articleB},  // Remove UserB's article
		}

		_, err := config.UserCollection.UpdateOne(context.TODO(), filterA, updateA)
		if err != nil {
			return nil, err
		}

		_, err = config.UserCollection.UpdateOne(context.TODO(), filterB, updateB)
		if err != nil {
			return nil, err
		}
	}

	return GetUserByID(userA)
} 