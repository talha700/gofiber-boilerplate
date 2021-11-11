package controllers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber-biolerplate/app/models"
	"github.com/gofiber-biolerplate/app/queries"
	"github.com/gofiber-biolerplate/pkg/utils"
)


// @Tags User
// @Summary Get Refresh Tokens
// @Description Get Refresh Tokens
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param Authorization header string true "Insert your access token" default(Bearer <token>)
// @Param data body models.Renew true "Give Refresh Token"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.RefreshTokenResponse
// @Router /token/refresh [post]
func RenewTokens(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current user.
	expiresAccessToken := claims.Expires

	// Checking, if now time greather than Access token expiration time.
	if now > expiresAccessToken {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create a new renew refresh token struct.
	renew := &models.Renew{}
	
	// Checking received data from JSON body.
	if err := c.BodyParser(renew); err != nil {
		// Return, if JSON data is not correct.
		log.Print(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from Refresh token of current user.
	expiresRefreshToken, err := utils.ParseRefreshToken(renew.RefreshToken)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if now time greather than Refresh token expiration time.
	if now < expiresRefreshToken {
		// Define user ID.
		userID := claims.UserID

		// Get user by ID.
		_, err := queries.GetUserByID(userID)
		if err != nil {
			// Return, if user not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "user with the given ID is not found",
			})
		}

		// Generate JWT Access & Refresh tokens.
		tokens, err := utils.GenerateNewTokens(userID)
		if err != nil {
			// Return status 500 and token generation error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"tokens": fiber.Map{
				"access":  tokens.Access,
				"refresh": tokens.Refresh,
			},
		})
	} else {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, your session was ended earlier",
		})
	}
}
