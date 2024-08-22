package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street    string    `json:"street" bson:"street"`
	City      string    `json:"city" bson:"city"`
	Postcode  int       `json:"postcode" bson:"postcode"`
	Citycode  int       `json:"citycode" bson:"citycode"`
	Floor     int       `json:"floor,omitempty" bson:"floor,omitempty"`
	Extra     string    `json:"extra,omitempty" bson:"extra,omitempty"`
	GeoPoints GeoPoints `json:"geopoints" bson:"geopoints"`
}

type GeoPoints struct {
	Type        string     `json:"type" bson:"type"`
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}

type ActivityStatus struct {
	LastConnected primitive.DateTime `json:"lastConnected" bson:"lastConnected"`
	Birthday      primitive.DateTime `json:"birthday" bson:"birthday"`
}

type BankInfo struct {
	IBAN string `json:"iban" bson:"iban"`
	BIC  string `json:"bic" bson:"bic"`
}

type User struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Version          int                `json:"version" bson:"version"`
	Pseudo           string             `json:"pseudo" bson:"pseudo"`
	Name             string             `json:"name" bson:"name"`
	Surname          string             `json:"surname" bson:"surname"`
	Address          Address            `json:"address" bson:"address"`
	Email            string             `json:"email" bson:"email"`
	Password         string             `json:"password" bson:"password"`
	Sexe             string             `json:"sexe" bson:"sexe"`
	PhoneNumber      string             `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	ActivityStatus   ActivityStatus     `json:"activityStatus" bson:"activityStatus"`
	BirthDate        time.Time          `json:"birthDate" bson:"birthDate"`
	BankInfo         *BankInfo          `json:"bankInfo,omitempty" bson:"bankInfo,omitempty"`
	AvatarUrl        string             `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	IsPremium        bool               `json:"isPremium" bson:"isPremium"`
	FavoriteArticles []string           `json:"favoriteArticles,omitempty" bson:"favoriteArticles,omitempty"`
	Credit           int                `json:"credit,omitempty" bson:"credit,omitempty"`
	Comments         []string           `json:"comments,omitempty" bson:"comments,omitempty"`
	Articles         []string           `json:"articles,omitempty" bson:"articles,omitempty"`
	Debit            []string           `json:"debit,omitempty" bson:"debit,omitempty"`
}
