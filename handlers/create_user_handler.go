package handlers

import (
	"fmt"
	"time"
	"trocup-user/models"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser handles the HTTP request to create a new user
func CreateUser(c *fiber.Ctx) error {
	// Create a new instance of the User model
	var user models.User

	// Parse the request body into the User model
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get Clerk-provided data (user id, name, surname, email) from the context
	clerkUserID := c.Locals("clerkUserId").(string)
	clerkEmail := c.Locals("clerkEmail").(string)
	clerkName := c.Locals("clerkName").(string)
	clerkSurname := c.Locals("clerkSurname").(string)

	// Assign Clerk-provided values to the user object
	user.ID = clerkUserID
	user.Email = clerkEmail
	user.Name = clerkName
	user.Surname = clerkSurname

	// Log the request body
	fmt.Printf("Request Body: %+v\n", user)

	// Set Birthday to the current time
	user.ActivityStatus.Birthday = primitive.NewDateTimeFromTime(time.Now())

	// Set LastConnected to the current time
	user.ActivityStatus.LastConnected = primitive.NewDateTimeFromTime(time.Now())

	// Check if the user already exists (by email or pseudo)
	err := services.CheckIfUserExists(c.Context(), user.Email, user.Pseudo)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create the new user (pass the context from Fiber)
	err = services.CreateUser(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
