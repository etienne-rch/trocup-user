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
	// Initialiser Fiber
	app := fiber.New()

	// Injecter le handler pour créer un utilisateur
	app.Post("/users", func(c *fiber.Ctx) error {
		// Simuler les données Clerk
		c.Locals("clerkUserId", "clerk_user_id_12345")
		c.Locals("clerkEmail", "john.doe@example.com")
		c.Locals("clerkName", "John")
		c.Locals("clerkSurname", "Doe")
		return handlers.CreateUser(c)
	})

	// Créer un utilisateur simulé
	user := models.User{
		Pseudo:      "testuser",
		Sexe:        "M",
		PhoneNumber: "1234567890",
		IsPremium:   false,
	}

	// Convertir l'utilisateur en JSON
	jsonUser, _ := json.Marshal(user)

	// Créer une requête POST simulée
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	// Exécuter la requête
	resp, _ := app.Test(req, -1)

	// Valider que la réponse a bien le statut HTTP 201 Created
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Nettoyage de la base de données après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
