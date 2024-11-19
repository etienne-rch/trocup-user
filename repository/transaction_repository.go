package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"
	"trocup-user/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUsersTransaction(articles []types.ArticleOwnership, isOneToOne bool) (map[string]*models.User, error) {
	updatedUsers := make(map[string]*models.User)
	
	if isOneToOne {
		// Handle 1-to-1 transaction
		article1 := articles[0]
		article2 := articles[1]
		
		// Update UserA (gives article1)
		updateA := bson.M{
			"$inc":  bson.M{"credit": -article1.Price},
			"$pull": bson.M{"articles": article1.ArticleID},
		}
		
		// Update UserB (gives article2)
		updateB := bson.M{
			"$inc":  bson.M{"credit": -article2.Price},
			"$pull": bson.M{"articles": article2.ArticleID},
		}

		if err := executeUpdates(article1.OwnerID, article2.OwnerID, updateA, updateB); err != nil {
			return nil, err
		}
	} else {
		// Handle 1-to-M transaction
		article := articles[0]
		
		// Update UserA (pays with balance)
		updateA := bson.M{
			"$inc": bson.M{"balance": -article.Price},
		}
		
		// Update UserB (gives article and loses credit)
		updateB := bson.M{
			"$inc": bson.M{"credit": -article.Price},
			"$pull": bson.M{"articles": article.ArticleID},
		}

		if err := executeUpdates(article.OwnerID, "", updateA, updateB); err != nil {
			return nil, err
		}
	}

	// Get updated users
	for _, article := range articles {
		user, err := GetUserByID(article.OwnerID)
		if err != nil {
			return nil, err
		}
		updatedUsers[article.OwnerID] = user
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