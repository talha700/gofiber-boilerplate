package controllers

import "github.com/gofiber/fiber/v2"





// @Tags Device
// @Summary Get Devices
// @Description Get Devices
// @Accept  json
// @Produce  json
// @Security JWT
// @Success 200 
// @Failure 400
// @Param Authorization header string true "Insert your access token" default(Bearer <token>)
// @Security BearerAuth
// @Router /device [get]
func GetDevice(c *fiber.Ctx) error {
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"device": 1,
	})
}
