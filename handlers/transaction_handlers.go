package handlers

import (
	"net/http"
	"os"
	"trocup-user/repository"

	"github.com/gofiber/fiber/v2"
)

type TransactionPayload struct {
	UserA      string   `json:"userA"`
	UserB      string   `json:"userB"`
	ArticleA   string   `json:"articleA"`  // Only for 1to1
	ArticleB   string   `json:"articleB,omitempty"`   
	ArticlePriceA float64 `json:"articleAPrice"` // Only for 1to1
	ArticlePriceB float64 `json:"articleBPrice,omitempty"` 
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

	// Perform a health check to the transaction micro-service
	transactionServiceURL := os.Getenv("TRANSACTION_SERVICE_URL")
	resp, err := http.Get(transactionServiceURL + "api/health")
	if err != nil || resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Transaction micro-service is unavailable",
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