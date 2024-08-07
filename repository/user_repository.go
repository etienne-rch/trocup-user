package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection

func InitUserRepository() {
    UserCollection = config.Client.Database("user_dev").Collection("user")
}

func CreateUser(user *models.User) error {
    _, err := UserCollection.InsertOne(context.TODO(), user)
    return err
}

func GetUsers() ([]models.User, error) {
    var users []models.User
    cursor, err := UserCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func GetUserByID(id primitive.ObjectID) (*models.User, error) {
    var user models.User
    filter := bson.M{"_id": id}
    err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}