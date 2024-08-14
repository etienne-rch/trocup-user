package handlers

import (
	"trocup-user/models"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/users", createUser)
    app.Get("/users", getUsers)
	app.Get("/users/:id", getUserByID)
}

func createUser(c *fiber.Ctx) error {
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    if err := services.CreateUser(user); err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(user)
}

func getUsers(c *fiber.Ctx) error {
    users, err := services.GetUsers()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(users)
}

func getUserByID(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(400).SendString("Invalid ID format")
    }

    user, err := services.GetUserByID(id)
    if err != nil {
        return c.Status(404).SendString("User not found")
    }

    return c.JSON(user)
}
