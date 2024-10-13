package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Label     string    `json:"label,omitempty" bson:"label,omitempty"`
	Street     string    `json:"street,omitempty" bson:"street,omitempty"`
	City       string    `json:"city,omitempty" bson:"city,omitempty"`
	Postcode   string       `json:"postcode,omitempty" bson:"postcode,omitempty"`
	Citycode   string       `json:"citycode,omitempty" bson:"citycode,omitempty"`
	Floor      int       `json:"floor,omitempty" bson:"floor,omitempty"`
	Extra      string    `json:"extra,omitempty" bson:"extra,omitempty"`
	GeoPoints  GeoPoints `json:"geopoints,omitempty" bson:"geopoints,omitempty"`
}

type GeoPoints struct {
	Type        string    `json:"type,omitempty" bson:"type,omitempty" validate:"eq=Point"`
	Coordinates []float64 `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
}

type ActivityStatus struct {
	LastConnected primitive.DateTime `json:"lastConnected,omitempty" bson:"lastConnected,omitempty"`
	Birthday      primitive.DateTime `json:"birthday,omitempty" bson:"birthday,omitempty"`
}

type BankInfo struct {
	IBAN string `json:"iban,omitempty" bson:"iban,omitempty"`
	BIC  string `json:"bic,omitempty" bson:"bic,omitempty"`
}

type User struct {
	ID               string             `json:"id,omitempty" bson:"_id,omitempty"` // Clerk's user ID
	Version          int                `json:"version,omitempty" bson:"version,omitempty"`
	Pseudo           string             `json:"pseudo" bson:"pseudo" validate:"required"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Surname          string             `json:"surname,omitempty" bson:"surname,omitempty" validate:"required"`
	Address          []Address            `json:"address,omitempty" bson:"address,omitempty"`
	Email            string             `json:"email" bson:"email,omitempty" validate:"required,email"`
	Sexe             string             `json:"sexe,omitempty" bson:"sexe,omitempty" validate:"oneof=M F"`
	PhoneNumber      string             `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty" validate:"omitempty,e164"`
	ActivityStatus   ActivityStatus     `json:"activityStatus,omitempty" bson:"activityStatus,omitempty"`
	BirthDate        time.Time          `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
	BankInfo         *BankInfo          `json:"bankInfo,omitempty" bson:"bankInfo,omitempty"`
	AvatarUrl        string             `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	IsPremium        bool               `json:"isPremium,omitempty" bson:"isPremium,omitempty"`
	FavoriteArticles []string           `json:"favoriteArticles,omitempty" bson:"favoriteArticles,omitempty"`
	Credit           int                `json:"credit,omitempty" bson:"credit,omitempty"`
	Comments         []string           `json:"comments,omitempty" bson:"comments,omitempty"`
	Articles         []string           `json:"articles,omitempty" bson:"articles,omitempty"`
	Debit            []string           `json:"debit,omitempty" bson:"debit,omitempty"`
}