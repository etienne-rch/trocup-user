package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck est une route simple pour vérifier l'état du service
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Service is up and running!",
	})
}
