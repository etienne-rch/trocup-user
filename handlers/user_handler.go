package handlers

import (
	"trocup-user/models"
	"trocup-user/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/users", createUser)
    app.Get("/users", getUsers)
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
