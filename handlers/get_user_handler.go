package handlers

import (
	"net/http"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
)

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")  // L'ID de Clerk est une chaîne de caractères

	// Appel du service avec l'ID string
	user, err := services.GetUserByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}
