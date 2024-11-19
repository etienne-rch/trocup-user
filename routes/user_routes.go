package routes

import (
	"fmt"
	"log"
	"trocup-user/handlers"
	"trocup-user/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// Add logging middleware for all requests
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request received: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	// Routes publiques : accessibles sans authentification
	public := app.Group("/api")

	public.Get("/health", handlers.HealthCheck)
	public.Get("/users/:id", handlers.GetUserByID)

	// Routes protégées : accessibles uniquement avec authentification
	protected := app.Group("/api/protected", middleware.ClerkAuthMiddleware)

	protected.Post("/users", handlers.CreateUser)
	protected.Put("/users/:id", handlers.UpdateUser)
	protected.Patch("/users/transactions", handlers.UpdateUsersTransaction)
	protected.Patch("/users/:id", handlers.UpdateUserArticle)
	protected.Delete("/users/:id", handlers.DeleteUser)

	// Routes accessibles uniquement aux utilisateurs connectés et admin : /api/protected/admin
	admin := protected.Group("/admin", middleware.ClerkAdminMiddleware)
	admin.Get("/users", handlers.GetUsers)

	// Ajouter une route catch-all pour le débogage
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString(fmt.Sprintf("Route not found: %s", c.Path()))
	})
}
