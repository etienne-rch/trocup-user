package tests

import (
	"testing"
	"time"
	"trocup-user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUser(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("failed to create ObjectID: %v", err)
	}

	now := time.Now()
	address := models.Address{
		Street:   "123 Main St",
		City:     "Anytown",
		Postcode: 12345,
		Citycode: 123,
		GeoPoints: models.GeoPoints{
			Type:        "Point",
			Coordinates: [2]bson.A{{1.0, 2.0}},
		},
	}

	activityStatus := models.ActivityStatus{
		LastConnected: primitive.NewDateTimeFromTime(now),
		Birthday:      primitive.NewDateTimeFromTime(now.AddDate(-30, 0, 0)),
	}

	user := models.User{
		ID:               id,
		Version:          1,
		Pseudo:           "johndoe",
		Name:             "John",
		Surname:          "Doe",
		Address:          address,
		Email:            "johndoe@example.com",
		Password:         "password",
		Sexe:             "M",
		ActivityStatus:   activityStatus,
		BirthDate:        now.AddDate(-30, 0, 0),
		IsPremium:        true,
		FavoriteArticles: []string{"article1", "article2"},
		Comments:         []string{"comment1", "comment2"},
		Articles:         []string{"article1", "article2"},
		Debit:            []string{"debit1", "debit2"},
	}

	if user.ID != id {
		t.Errorf("expected ID to be %v, got %v", id, user.ID)
	}
	if user.Version != 1 {
		t.Errorf("expected Version to be 1, got %d", user.Version)
	}
	if user.Pseudo != "johndoe" {
		t.Errorf("expected Pseudo to be 'johndoe', got %s", user.Pseudo)
	}
	if user.Name != "John" {
		t.Errorf("expected Name to be 'John', got %s", user.Name)
	}
	if user.Surname != "Doe" {
		t.Errorf("expected Surname to be 'Doe', got %s", user.Surname)
	}
	// Compare Address fields
	if user.Address.Street != address.Street {
		t.Errorf("expected Street to be %s, got %s", address.Street, user.Address.Street)
	}
	if user.Address.City != address.City {
		t.Errorf("expected City to be %s, got %s", address.City, user.Address.City)
	}
	if user.Address.Postcode != address.Postcode {
		t.Errorf("expected Postcode to be %d, got %d", address.Postcode, user.Address.Postcode)
	}
	if user.Address.Citycode != address.Citycode {
		t.Errorf("expected Citycode to be %d, got %d", address.Citycode, user.Address.Citycode)
	}
	if user.Address.GeoPoints.Type != address.GeoPoints.Type {
		t.Errorf("expected GeoPoints.Type to be %s, got %s", address.GeoPoints.Type, user.Address.GeoPoints.Type)
	}
	if user.Address.GeoPoints.Coordinates[0][0] != address.GeoPoints.Coordinates[0][0] || user.Address.GeoPoints.Coordinates[0][1] != address.GeoPoints.Coordinates[0][1] {
		t.Errorf("expected GeoPoints.Coordinates to be %v, got %v", address.GeoPoints.Coordinates, user.Address.GeoPoints.Coordinates)
	}

	if user.Email != "johndoe@example.com" {
		t.Errorf("expected Email to be 'johndoe@example.com', got %s", user.Email)
	}
	if user.Password != "password" {
		t.Errorf("expected Password to be 'password', got %s", user.Password)
	}
	if user.Sexe != "M" {
		t.Errorf("expected Sexe to be 'M', got %s", user.Sexe)
	}
	if user.ActivityStatus != activityStatus {
		t.Errorf("expected ActivityStatus to be %v, got %v", activityStatus, user.ActivityStatus)
	}
	if !user.BirthDate.Equal(now.AddDate(-30, 0, 0)) {
		t.Errorf("expected BirthDate to be %v, got %v", now.AddDate(-30, 0, 0), user.BirthDate)
	}
	if user.IsPremium != true {
		t.Errorf("expected IsPremium to be true, got %v", user.IsPremium)
	}
	if len(user.FavoriteArticles) != 2 || user.FavoriteArticles[0] != "article1" || user.FavoriteArticles[1] != "article2" {
		t.Errorf("expected FavoriteArticles to be ['article1', 'article2'], got %v", user.FavoriteArticles)
	}
	if len(user.Comments) != 2 || user.Comments[0] != "comment1" || user.Comments[1] != "comment2" {
		t.Errorf("expected Comments to be ['comment1', 'comment2'], got %v", user.Comments)
	}
	if len(user.Articles) != 2 || user.Articles[0] != "article1" || user.Articles[1] != "article2" {
		t.Errorf("expected Articles to be ['article1', 'article2'], got %v", user.Articles)
	}
	if len(user.Debit) != 2 || user.Debit[0] != "debit1" || user.Debit[1] != "debit2" {
		t.Errorf("expected Debit to be ['debit1', 'debit2'], got %v", user.Debit)
	}
}
