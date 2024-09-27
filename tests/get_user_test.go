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
)

func TestGetUserByID(t *testing.T) {
	app := fiber.New()

	app.Get("/users/:id", handlers.GetUserByID)

	user := models.User{
		ID:          "clerk_user_id_12345", // Utilisation du Clerk ID
		Version:     1,
		Pseudo:      "testuser",
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		Sexe:        "M",
		PhoneNumber: "1234567890",
		IsPremium:   false,
	}

	// Insertion de l'utilisateur dans la base de données
	config.UserCollection.InsertOne(context.TODO(), user)

	// Effectuer la requête GET avec le Clerk ID
	req := httptest.NewRequest("GET", "/users/"+user.ID, nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}

