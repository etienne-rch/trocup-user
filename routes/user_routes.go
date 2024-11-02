package routes

import (
	"trocup-user/handlers"
	"trocup-user/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// Routes publiques : accessibles sans authentification
	public := app.Group("/api")
	
	public.Get("/health", handlers.HealthCheck)
	public.Get("/users/:id", handlers.GetUserByID)

	// Routes protégées : accessibles uniquement avec authentification
	protected := app.Group("/api/protected", middleware.ClerkAuthMiddleware)

	protected.Post("/users", handlers.CreateUser)
	protected.Get("/users", handlers.GetUsers)
	protected.Put("/users/:id", handlers.UpdateUser)
	protected.Delete("/users/:id", handlers.DeleteUser)
}
