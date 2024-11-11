package handlers

import (
	"trocup-user/repository"

	"github.com/gofiber/fiber/v2"
)

// UpdateUserArticle handles credit and article updates from the article service
func UpdateUserArticle(c *fiber.Ctx) error {
    userID := c.Params("id")

    var articlePayload struct {
        ArticleID string  `json:"articleId"`
        Price     float64 `json:"price"`
    }
    
    if err := c.BodyParser(&articlePayload); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid credit update payload",
            "details": err.Error(),
        })
    }

    // First check if user exists
    exists, err := repository.UserExists(userID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error checking user existence",
            "details": err.Error(),
        })
    }

    if !exists {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
            "userId": userID,
        })
    }

    updatedUser, err := repository.UpdateUserArticleCredit(userID, articlePayload.ArticleID, articlePayload.Price)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update user credit",
            "details": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(updatedUser)
} 