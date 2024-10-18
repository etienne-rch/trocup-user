package handlers

import (
	"os"
	"strconv"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	// Get default skip and limit from environment variables
	defaultSkip, _ := strconv.ParseInt(os.Getenv("DEFAULT_SKIP"), 10, 64)
	defaultLimit, _ := strconv.ParseInt(os.Getenv("DEFAULT_LIMIT"), 10, 64)

	// If env variables aren't set or parsing fails, fallback to hardcoded defaults
	if defaultSkip < 0 {
		defaultSkip = 0
	}
	if defaultLimit <= 0 {
		defaultLimit = 100
	}

	// Get query parameters with defaults from env variables
	skipParam := c.Query("skip", strconv.FormatInt(defaultSkip, 10))    // Default skip from env
	limitParam := c.Query("limit", strconv.FormatInt(defaultLimit, 10)) // Default limit from env

	skip, err := strconv.ParseInt(skipParam, 10, 64)
	if err != nil || skip < 0 {
		skip = defaultSkip
	}

	limit, err := strconv.ParseInt(limitParam, 10, 64)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	users, err := services.GetUsers(skip, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"skip":  skip,
		"limit": limit,
		"users": users,
	})
}
