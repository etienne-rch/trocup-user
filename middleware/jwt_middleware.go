package middleware

import (
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

	
	// Store Clerk user data in Fiber's context for future handlers
	c.Locals("clerkUserId", claims.Subject)
	c.Locals("clerkEmail", usr.EmailAddresses[0].EmailAddress)
	c.Locals("clerkName", firstName)
	c.Locals("clerkSurname", lastName)

	// Continue to the next handler in the chain
	return c.Next()
}