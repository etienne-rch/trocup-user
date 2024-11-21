package handlers

import (
	"log"
	"trocup-user/repository"
	"trocup-user/types"

	"github.com/gofiber/fiber/v2"
)

type TransactionPayload struct {
	UserA         string  `json:"userA"`
	UserB         string  `json:"userB"`
	ArticleA      string  `json:"articleA,omitempty"` // Only for 1to1
	ArticleB      string  `json:"articleB"`
	ArticlePriceA float64 `json:"articlePriceA,omitempty"` // Only for 1to1
	ArticlePriceB float64 `json:"articlePriceB"`
}

func UpdateUsersTransaction(c *fiber.Ctx) error {
	var payload TransactionPayload

	if err := c.BodyParser(&payload); err != nil {
		log.Printf("Error parsing payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid payload",
			"details": err.Error(),
		})
	}

	log.Printf("Received payload: %+v", payload)
	log.Printf("Request Origin: %s", c.Get("Origin"))

	// Common validation for both types of transactions
	if payload.UserA == "" || payload.UserB == "" ||
		payload.ArticleB == "" || payload.ArticlePriceB <= 0 {
		log.Printf("Missing required fields - UserA: %s, UserB: %s, ArticleB: %s, ArticlePriceB: %f",
			payload.UserA, payload.UserB, payload.ArticleB, payload.ArticlePriceB)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Missing required fields",
			"details": "userA, userB, articleB, and articlePriceB are required",
		})
	}

	// Determine transaction type and validate
	isOneToOne := payload.ArticleA != "" || payload.ArticlePriceA > 0
	var articles []types.ArticleOwnership

	log.Printf("isOneToOne: %t", isOneToOne)

	if isOneToOne {
		if payload.ArticleA == "" || payload.ArticlePriceA <= 0 {
			log.Printf("Incomplete 1-to-1 transaction data - ArticleA: %s, ArticlePriceA: %f",
				payload.ArticleA, payload.ArticlePriceA)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid 1-to-1 transaction",
				"details": "Both articleA and articlePriceA must be provided for 1-to-1 transactions",
			})
		}
		articles = []types.ArticleOwnership{
			{ArticleID: payload.ArticleA, UserID: payload.UserA, Price: payload.ArticlePriceA},
			{ArticleID: payload.ArticleB, UserID: payload.UserB, Price: payload.ArticlePriceB},
		}

	} else {
		log.Printf("🔥 1-to-M transaction - in handler")

		articles = []types.ArticleOwnership{
			{ArticleID: "", UserID: payload.UserA, Price: 0},
			{ArticleID: payload.ArticleB, UserID: payload.UserB, Price: payload.ArticlePriceB},
		}
	}

	updatedUsers, err := repository.UpdateUsersTransaction(articles, isOneToOne)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUsers)
}
