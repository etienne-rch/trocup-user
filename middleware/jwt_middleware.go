package middleware

import (
	"log"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
)

// ClerkAuthMiddleware verifies the Authorization header and Clerk session
func ClerkAuthMiddleware(c *fiber.Ctx) error {
	// Extract the Authorization header from the request
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		log.Println("Missing Authorization header")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	// Extract the Bearer token from the Authorization header
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		log.Println("Invalid Authorization format")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization format",
		})
	}

	// Log the token for debugging (optional, avoid this in production with sensitive data)
	log.Printf("Received Token: %s\n", token)

	// Verify the session token using Clerk's SDK
	claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{Token: token})
	if err != nil {
		log.Println("Invalid or expired token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Log claims (subject) for debugging purposes
	log.Printf("Session Claims - Subject (UserID): %s\n", claims.Subject)

	// Fetch the user associated with the session token
	usr, err := user.Get(c.Context(), claims.Subject)
	if err != nil {
		log.Printf("Failed to retrieve user: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user information",
		})
	}

	// Dereference pointers to get actual string values (check for nil pointers)
	firstName := ""
	lastName := ""
	if usr.FirstName != nil {
		firstName = *usr.FirstName
	}
	if usr.LastName != nil {
		lastName = *usr.LastName
	}

	// Log the user info retrieved from Clerk for debugging purposes
	log.Printf("User Info - ID: %s, Name: %s %s, Email: %s, Banned: %v\n", usr.ID, firstName, lastName, usr.EmailAddresses[0].EmailAddress, usr.Banned)

	// Store Clerk user data in Fiber's context for future handlers
	c.Locals("clerkUserId", claims.Subject)
	c.Locals("clerkEmail", usr.EmailAddresses[0].EmailAddress)
	c.Locals("clerkName", firstName)
	c.Locals("clerkSurname", lastName)

	// Continue to the next handler in the chain
	return c.Next()
}