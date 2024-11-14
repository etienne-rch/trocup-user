package handlers

import (
	"trocup-user/repository"

	"github.com/gofiber/fiber/v2"
)

type TransactionPayload struct {
	UserA      string   `json:"userA"`
	UserB      string   `json:"userB"`
	ArticleA   string   `json:"articleA,omitempty"`  // Only for 1to1
	ArticleB   string   `json:"articleB"`   
	ArticlePriceA float64 `json:"articlePriceA,omitempty"` // Only for 1to1
	ArticlePriceB float64 `json:"articlePriceB"` 
}

func UpdateUsersTransaction(c *fiber.Ctx) error {
	var payload TransactionPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	// Validate required fields
	if payload.UserA == "" || payload.UserB == "" || 
	   payload.ArticleA == "" || payload.ArticlePriceA <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}


	// Check if it's a 1to1 transaction
	isOneToOne := payload.ArticleB != "" && payload.ArticlePriceB > 0

	updatedUser, err := repository.UpdateUsersTransaction(
		payload.UserA, 
		payload.UserB,
		payload.ArticleA,
		payload.ArticleB,
		payload.ArticlePriceA,
		payload.ArticlePriceB,
		isOneToOne,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
} 	
