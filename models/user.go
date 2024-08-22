package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
    Street   string `json:"street" bson:"street" validate:"required"`
    City     string `json:"city" bson:"city" validate:"required"`
    Postcode int    `json:"postcode" bson:"postcode" validate:"required"`
    Citycode int    `json:"citycode" bson:"citycode"`
    Floor    int    `json:"floor,omitempty" bson:"floor,omitempty"`
    Extra    string `json:"extra,omitempty" bson:"extra,omitempty"`
    GeoPoints  GeoPoints     `json:"geopoints" bson:"geopoints" validate:"required"`
}

type GeoPoints struct {
    Type        string           `json:"type" bson:"type" validate:"required,eq=Point"`
    Coordinates []float64 `json:"coordinates" bson:"coordinates" validate:"required"`
}

type ActivityStatus struct {
    LastConnected primitive.DateTime `json:"lastConnected" bson:"lastConnected"`
    Birthday      primitive.DateTime `json:"birthday" bson:"birthday"`
}

type BankInfo struct {
    IBAN string `json:"iban" bson:"iban" validate:"required"`
    BIC  string `json:"bic" bson:"bic" validate:"required"`
}

type User struct {
    ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Version          int                `json:"version" bson:"version"`
    Pseudo           string             `json:"pseudo" bson:"pseudo" validate:"required"`
    Name             string             `json:"name" bson:"name" validate:"required"`
    Surname          string             `json:"surname" bson:"surname" validate:"required"`
    Address          Address            `json:"address" bson:"address" validate:"required,dive"`
    Email            string             `json:"email" bson:"email" validate:"required,email"`
    Password         string             `json:"password" bson:"password" validate:"required,min=6"`
    Sexe             string             `json:"sexe" bson:"sexe" validate:"required,oneof=M F"`
    PhoneNumber      string             `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty" validate:"omitempty,e164"`
    ActivityStatus   ActivityStatus     `json:"activityStatus" bson:"activityStatus"`
    BirthDate        time.Time          `json:"birthDate" bson:"birthDate" validate:"required"`
    BankInfo         *BankInfo          `json:"bankInfo,omitempty" bson:"bankInfo,omitempty"`
    AvatarUrl        string             `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
    IsPremium        bool               `json:"isPremium" bson:"isPremium"`
    FavoriteArticles []string           `json:"favoriteArticles,omitempty" bson:"favoriteArticles,omitempty"`
    Credit           int                `json:"credit,omitempty" bson:"credit,omitempty"`
    Comments         []string           `json:"comments,omitempty" bson:"comments,omitempty"`
    Articles         []string           `json:"articles,omitempty" bson:"articles,omitempty"`
    Debit            []string           `json:"debit,omitempty" bson:"debit,omitempty"`
}
