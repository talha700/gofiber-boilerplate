package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber-biolerplate/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/user/login", controllers.Login) // auth, return Access & Refresh tokens
	route.Post("/user/signup", controllers.CreateUser) // auth, return Access & Refresh tokens
	


}
