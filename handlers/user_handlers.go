package handlers

import (
	"trocup-user/models"
	"trocup-user/repository"

	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
    userID := c.Params("id")
    
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user update payload",
        })
    }

    updatedUser, err := repository.UpdateUser(userID, &user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update user",
        })
    }

    return c.Status(fiber.StatusOK).JSON(updatedUser)
}