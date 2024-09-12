package tests

import (
	"bytes"
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

func TestCreateUser(t *testing.T) {
	app := fiber.New()

	app.Post("/users", handlers.CreateUser)

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

	jsonUser, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
