package repository

import (
	"context"
	"fmt"
	"log"
	"trocup-user/config"
	"trocup-user/models"
	"trocup-user/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUsersTransaction(articles []types.ArticleOwnership, isOneToOne bool) (map[string]*models.User, error) {

	log.Printf("Articles: %+v", articles)
	log.Printf("IsOneToOne: %t", isOneToOne)

	userAparam := articles[0]
	userBparam := articles[1]

	// Get users data from DB
	userA, err := GetUserByID(userAparam.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get userA: %v", err)
	}

	_, err = GetUserByID(userBparam.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get userB: %v", err)
	}

	if !isOneToOne {
		// For 1-to-M, check userA's balance and premium status

		// Calculate minimum allowed balance
		minAllowedBalance := float64(0)
		if userA.IsPremium {
			minAllowedBalance = -userA.Credit
		} else {
			minAllowedBalance = 0
		}

		// Check if transaction would exceed minimum allowed balance
		if (userA.Balance - userBparam.Price) < minAllowedBalance {
			log.Printf("❌ Insufficient balance: %.2f (min allowed: %.2f)", userA.Balance, minAllowedBalance)
			return nil, fmt.Errorf("❌ Insufficient balance: %.2f (min allowed: %.2f)", userA.Balance, minAllowedBalance)
		}
	}

	// Proceed with updates if validation passed
	if isOneToOne {

		updateA := bson.M{
			"$inc":  bson.M{"credit": -userAparam.Price},
			"$pull": bson.M{"articles": userAparam.ArticleID},
		}

		updateB := bson.M{
			"$inc":  bson.M{"credit": -userBparam.Price},
			"$pull": bson.M{"articles": userBparam.ArticleID},
		}

		if err := executeUpdates(userAparam.UserID, userBparam.UserID, updateA, updateB); err != nil {
			return nil, err
		}
	} else {
		// 1-to-M transaction

		updateA := bson.M{
			"$inc": bson.M{"balance": -userBparam.Price},
		}

		updateB := bson.M{
			"$inc": bson.M{
				"credit":  -userBparam.Price,
				"balance": userBparam.Price,
			},
			"$pull": bson.M{"articles": userBparam.ArticleID},
		}

		if err := executeUpdates(userAparam.UserID, userBparam.UserID, updateA, updateB); err != nil {
			return nil, err
		}
	}

	// Fetch both updated users after the transaction
	updatedUserA, err := GetUserByID(userAparam.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated userA: %v", err)
	}

	updatedUserB, err := GetUserByID(userBparam.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated userB: %v", err)
	}

	// Return both updated users
	updatedUsers := map[string]*models.User{
		userAparam.UserID: updatedUserA,
		userBparam.UserID: updatedUserB,
	}
	return updatedUsers, nil
}

func executeUpdates(userA, userB string, updateA, updateB bson.M) error {
	session, err := config.UserCollection.Database().Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.TODO())

	_, err = session.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		if _, err := config.UserCollection.UpdateOne(ctx, bson.M{"_id": userA}, updateA); err != nil {
			return nil, err
		}

		if userB != "" {
			if _, err := config.UserCollection.UpdateOne(ctx, bson.M{"_id": userB}, updateB); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}
