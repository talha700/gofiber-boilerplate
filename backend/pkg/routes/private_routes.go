package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber-biolerplate/app/controllers"
	"github.com/gofiber-biolerplate/pkg/middleware"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/device", middleware.JWTProtected(), controllers.GetDevice)
	route.Post("token/refresh", middleware.JWTProtected(), controllers.RenewTokens)
	route.Patch("/user/changepass", controllers.ChangePassword)

}
