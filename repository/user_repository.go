package repository

import (
	"context"
	"trocup-user/config"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserRepository() {
    userCollection = config.Client.Database("user_dev").Collection("user")
}

func CreateUser(user *models.User) error {
    _, err := userCollection.InsertOne(context.TODO(), user)
    return err
}

func GetUsers() ([]models.User, error) {
    var users []models.User
    cursor, err := userCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }
    return users, nil
}
