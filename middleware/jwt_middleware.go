package middleware

import (
	"encoding/json"
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

	// Decode PrivateMetadata
	var privateMetadata map[string]interface{}
	if err := json.Unmarshal(usr.PrivateMetadata, &privateMetadata); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode private metadata",
		})
	}

	// Store Clerk user data and metadata in Fiber's context
	c.Locals("clerkUserId", claims.Subject)
	c.Locals("clerkEmail", usr.EmailAddresses[0].EmailAddress)
	c.Locals("clerkName", firstName)
	c.Locals("clerkSurname", lastName)
	c.Locals("clerkPrivateMetadata", privateMetadata) // Store private metadata here

	// Continue to the next handler in the chain
	return c.Next()
}

// ClerkAdminMiddleware checks if the user has an "admin" role
func ClerkAdminMiddleware(c *fiber.Ctx) error {
	// Retrieve private metadata from Fiber's context (set in ClerkAuthMiddleware)
	privateMetadata, ok := c.Locals("clerkPrivateMetadata").(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Unable to access private metadata",
		})
	}

	// Check if "role" exists and if it is "admin"
	role, ok := privateMetadata["role"].(string)
	if !ok || role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied: Admins only",
		})
	}

	// Continue to the next handler if the user is an admin
	return c.Next()
}
