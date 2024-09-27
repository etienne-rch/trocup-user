package handlers

import (
	"net/http"
	"trocup-user/models"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")  // Utiliser directement l'ID string de Clerk

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedUser, err := services.UpdateUser(id, &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedUser)
}
