package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trocup-user/config"
	"trocup-user/handlers"
	"trocup-user/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	app := fiber.New()

	app.Put("/users/:id", handlers.UpdateUser)

	// Insérer un utilisateur pour le test
	user := models.User{
		ID:          "clerk_user_id_12345",
		Version:     1,
		Pseudo:      "testuser",
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		Sexe:        "M",
		PhoneNumber: "1234567890",
		IsPremium:   false,
	}
	config.UserCollection.InsertOne(context.TODO(), user)

	// Modifier les données de l'utilisateur
	updatedUser := models.User{
		Pseudo:      "updateduser",
		Name:        "Jane",
		Surname:     "Doe",
		Email:       "jane.doe@example.com",
		Sexe:        "F",
		PhoneNumber: "0987654321",
		IsPremium:   true,
	}

	jsonUser, _ := json.Marshal(updatedUser)
	req := httptest.NewRequest("PUT", "/users/"+user.ID, bytes.NewReader(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
