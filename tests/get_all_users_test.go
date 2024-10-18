package tests

import (
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

func TestGetUsers(t *testing.T) {
	app := fiber.New()

	app.Get("/users", handlers.GetUsers)

	users := []models.User{
		{
			ID:          "clerk_user_id_12345",
			Version:     1,
			Pseudo:      "testuser1",
			Name:        "John",
			Surname:     "Doe",
			Email:       "john.doe1@example.com",
			Sexe:        "M",
			PhoneNumber: "1234567890",
			IsPremium:   false,
		},
		{
			ID:          "clerk_user_id_12345",
			Version:     1,
			Pseudo:      "testuser2",
			Name:        "Jane",
			Surname:     "Doe",
			Email:       "jane.doe@example.com",
			Sexe:        "F",
			PhoneNumber: "0987654321",
			IsPremium:   true,
		},
	}

	// Insérer les utilisateurs créés dans la base de données de test
	for _, user := range users {
		config.UserCollection.InsertOne(context.TODO(), user)
	}

	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err, "Failed to get response from server")

	// Vérifier que la réponse HTTP renvoie un statut 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Décoder le JSON de la réponse pour obtenir les utilisateurs
	var response struct {
		Users []models.User `json:"users"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err, "Failed to decode response body")

	// Vérifier que le nombre d'utilisateurs retournés correspond à ceux que j'ai insérés dans la base de données
	assert.Equal(t, len(users), len(response.Users), "Number of returned users does not match expected count")

	// Comparer les champs des utilisateurs retournés avec ceux que j'ai insérés
	for i, user := range users {
		assert.Equal(t, user.Pseudo, response.Users[i].Pseudo)
		assert.Equal(t, user.Name, response.Users[i].Name)
		assert.Equal(t, user.Surname, response.Users[i].Surname)
		assert.Equal(t, user.Email, response.Users[i].Email)
		assert.Equal(t, user.Sexe, response.Users[i].Sexe)
		assert.Equal(t, user.PhoneNumber, response.Users[i].PhoneNumber)
		assert.Equal(t, user.IsPremium, response.Users[i].IsPremium)
	}

	// Nettoyage après chaque test
	defer config.CleanUpTestDatabase("test_db")
}
