package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"trocup-user/config"
	"trocup-user/handlers"
	"trocup-user/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUsers(t *testing.T) {
	app := fiber.New()

	app.Get("/users", handlers.GetUsers)

	users := []models.User{
		{
			ID:          primitive.NewObjectID(),
			Version:     1,
			Pseudo:      "testuser1",
			Name:        "John",
			Surname:     "Doe",
			Email:       "john.doe1@example.com",
			Password:    "password123",
			Sexe:        "M",
			PhoneNumber: "1234567890",
			IsPremium:   false,
		},
		{
			ID:          primitive.NewObjectID(),
			Version:     1,
			Pseudo:      "testuser2",
			Name:        "Jane",
			Surname:     "Doe",
			Email:       "jane.doe@example.com",
			Password:    "password123",
			Sexe:        "F",
			PhoneNumber: "0987654321",
			IsPremium:   true,
		},
	}
	for _, user := range users {
		config.UserCollection.InsertOne(context.TODO(), user)
	}

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}