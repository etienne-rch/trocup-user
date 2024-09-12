package handlers

import (
	"log"
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
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Print the parsed body from the request
	log.Println("Parsed User from Body:")
	log.Printf("Pseudo: %s, AvatarUrl: %s, ActivityStatus.Birthday: %v\n", user.Pseudo, user.AvatarUrl, user.ActivityStatus.Birthday)

	// Get Clerk-provided data (name, surname, email) from the context
	clerkUserID := c.Locals("clerkUserId").(string)
	clerkEmail := c.Locals("clerkEmail").(string)
	clerkName := c.Locals("clerkName").(string)
	clerkSurname := c.Locals("clerkSurname").(string)

	// Print the information from Clerk context
	log.Println("Clerk Info from Context:")
	log.Printf("UserID: %s, Email: %s, Name: %s, Surname: %s\n", clerkUserID, clerkEmail, clerkName, clerkSurname)

	// Assign Clerk-provided values to the user object
	user.ID = clerkUserID
	user.Email = clerkEmail
	user.Name = clerkName
	user.Surname = clerkSurname
	user.IsPremium = false // Default value for isPremium

	// Set Birthday to the current time
	user.ActivityStatus.Birthday = primitive.NewDateTimeFromTime(time.Now())

	// Set LastConnected to the current time
	user.ActivityStatus.LastConnected = primitive.NewDateTimeFromTime(time.Now())

	// Print the final user object before creating the user in the database
	log.Println("Final User Object:")
	log.Printf("ID: %s, Name: %s, Email: %s, Pseudo: %s, LastConnected: %v\n", user.ID, user.Name, user.Email, user.Pseudo, user.ActivityStatus.LastConnected)

	// Check if the user already exists (by email or pseudo)
	err := services.CheckIfUserExists(c.Context(), user.Email, user.Pseudo)
	if err != nil {
		log.Println("User already exists with Email or Pseudo:", err)
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create the new user (pass the context from Fiber)
	err = services.CreateUser(c.Context(), &user)
	if err != nil {
		log.Println("Failed to create user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Return success response
	log.Println("User created successfully")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}