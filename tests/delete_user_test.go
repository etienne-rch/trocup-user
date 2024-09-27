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

func TestDeleteUser(t *testing.T) {
	app := fiber.New()

	app.Delete("/users/:id", handlers.DeleteUser)

	user := models.User{
		ID:          "clerk_user_id_12345", // Utilisation d'un ID string
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

	req := httptest.NewRequest("DELETE", "/users/"+user.ID, nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}

