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

func TestGetUserByID(t *testing.T) {
	app := fiber.New()

	app.Get("/users/:id", handlers.GetUserByID)

	user := models.User{
		ID:          primitive.NewObjectID(),
		Version:     1,
		Pseudo:      "testuser",
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		Password:    "password123",
		Sexe:        "M",
		PhoneNumber: "1234567890",
		IsPremium:   false,
	}
	config.UserCollection.InsertOne(context.TODO(), user)

	req := httptest.NewRequest("GET", "/users/"+user.ID.Hex(), nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
