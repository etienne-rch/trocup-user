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
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	// Extract the Bearer token from the Authorization header
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization format",
		})
	}

	// Verify the session token using Clerk's SDK
	claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{Token: token})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Fetch the user associated with the session token
	usr, err := user.Get(c.Context(), claims.Subject)
	if err != nil {
		log.Printf("Failed to retrieve user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user information",
		})
	}

	// Log user info for debugging purposes
	log.Printf("User: %s, Banned: %v", usr.ID, usr.Banned)

	// Store the Clerk user ID in Fiber's context for future handlers
	c.Locals("clerkUserId", claims.Subject)

	// Continue to the next handler in the chain
	return c.Next()
}