package tests

import (
	"reflect"
	"testing"
	"time"
	"trocup-user/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUser(t *testing.T) {
	id := "clerk_user_id_12345"
	

	now := time.Now()
	address := []models.Address{{
		Street:    "123 Test St",
		City:      "Test City",
		Postcode:  "12345",
		Citycode:  "123",
		GeoPoints: models.GeoPoints{
			Type:        "Point",
			Coordinates: []float64{1.23, 4.56},
		},
	}}

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
		Sexe:             "M",
		ActivityStatus:   activityStatus,
		BirthDate:        now.AddDate(-30, 0, 0),
		IsPremium:        true,
		FavoriteArticles: []string{"article1", "article2"},
		Comments:         []string{"comment1", "comment2"},
		Articles:         []string{"article1", "article2"},
		Credit:           100,
		Balance:          100,
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
	assert.Equal(t, "123 Test St", user.Address[0].Street, "Street should match")
	assert.Equal(t, "123 Test St", user.Address[0].Street, "Street in Address should match")

	assert.Equal(t, "Test City", user.Address[0].City, "City should match")
	assert.Equal(t, "Test City", user.Address[0].City, "City in Address should match")

	assert.NotNil(t, user.Address[0].GeoPoints, "GeoPoints should not be nil")
	assert.Equal(t, "Point", user.Address[0].GeoPoints.Type, "GeoPoints Type should be 'Point'")

	assert.Len(t, user.Address[0].GeoPoints.Coordinates, 2, "GeoPoints should have 2 coordinates")
	assert.Equal(t, 1.23, user.Address[0].GeoPoints.Coordinates[0], "Latitude should match")
	assert.Equal(t, 4.56, user.Address[0].GeoPoints.Coordinates[1], "Longitude should match")

	if user.Email != "johndoe@example.com" {
		t.Errorf("expected Email to be 'johndoe@example.com', got %s", user.Email)
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
	if !reflect.DeepEqual(user.FavoriteArticles, []string{"article1", "article2"}) {
		t.Errorf("expected FavoriteArticles to be ['article1', 'article2'], got %v", user.FavoriteArticles)
	}	
	if len(user.Comments) != 2 || user.Comments[0] != "comment1" || user.Comments[1] != "comment2" {
		t.Errorf("expected Comments to be ['comment1', 'comment2'], got %v", user.Comments)
	}
	if len(user.Articles) != 2 || user.Articles[0] != "article1" || user.Articles[1] != "article2" {
		t.Errorf("expected Articles to be ['article1', 'article2'], got %v", user.Articles)
	}
	if user.Credit != 100 {
		t.Errorf("expected Credit to be 100, got %f", user.Credit)
	}
	if user.Balance != 100 {
		t.Errorf("expected Balance to be 100, got %f", user.Balance)
	}
}