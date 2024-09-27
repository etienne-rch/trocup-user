package handlers

import (
	"net/http"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")  // L'ID de Clerk est récupéré sous forme de chaîne

	// Appel du service avec l'ID string
	if err := services.DeleteUser(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}

