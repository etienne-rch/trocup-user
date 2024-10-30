package routes

import (
	"trocup-user/handlers"
	"trocup-user/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/health", handlers.HealthCheck)

	// Routes publiques : accessibles sans authentification
	app.Get("/users/:id", handlers.GetUserByID)

	//app.Post("/register", handlers.Register)
	//app.Post("/login", handlers.Login)

	api := app.Group("/api", middleware.ClerkAuthMiddleware)

	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetUsers)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)
}
