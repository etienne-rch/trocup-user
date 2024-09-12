package handlers

import (
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
)

func GetUserByID(c *fiber.Ctx) error {
	// Get user ID from params
	userID := c.Params("id")

	// Get context from request
	ctx := c.Context()

	// Fetch user details from Clerk using the user ID
	usr, err := user.Get(ctx, userID)
	if err != nil {
		// Handle the error, return 500 if something went wrong
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// If the user is not found
	if usr == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Return user details in the response
	return c.JSON(fiber.Map{
		"id":        usr.ID,
		"firstName": *usr.FirstName,
		"lastName":  *usr.LastName,
		"email":     *usr.Email,
	})
}