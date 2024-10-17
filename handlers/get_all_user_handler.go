package handlers

import (
	"strconv"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	skipParam := c.Query("skip", "0")    // default value: skip = 0
	limitParam := c.Query("limit", "10") // default value: limit = 10

	skip, err := strconv.ParseInt(skipParam, 10, 64)
	if err != nil || skip < 0 {
		skip = 0
	}

	limit, err := strconv.ParseInt(limitParam, 10, 64)
	if err != nil || limit <= 0 {
		limit = 10
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
