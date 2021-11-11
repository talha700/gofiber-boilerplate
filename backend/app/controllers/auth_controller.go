package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber-biolerplate/app/models"
	"github.com/gofiber-biolerplate/app/queries"
	"github.com/gofiber-biolerplate/pkg/utils"
	"github.com/gofiber-biolerplate/platform/database"
)



// @Tags User
// @Summary User Login
// @Description Get Access and Refresh Tokens
// @Accept  json
// @Produce  json
// @Param data body models.SignIn true "Give username and Password"
// @Success 200 {object} models.LoginResponse
// @Failure 400
// @Router /user/login [post]
func Login(c *fiber.Ctx) error {
	// Create a new user auth struct.
	signIn := &models.SignIn{}

	// Checking received data from JSON body.
	if err := c.BodyParser(signIn); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.Connection()
	// Get user by email.
	foundedUser, err := queries.GetUserByUserName(db, signIn)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email is not found",
		})
	}

	// Compare given user password with stored in found user.
	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	if !compareUserPassword {
		// Return, if password is not compare to stored in database.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "wrong user email address or password",
		})
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(strconv.FormatUint(uint64(foundedUser.ID), 10))
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}




// @Tags User
// @Summary Change User Password
// @Description change user password
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param Authorization header string true "Insert your access token" default(Bearer <token>)
// @Param data body models.EditPassword true "Give old and new password"
// @Success 200 {object} models.ChangePass
// @Failure 400
// @Router /user/changepass [post]
func ChangePassword(c *fiber.Ctx) error {

	passwords := &models.EditPassword{}

	if err := c.BodyParser(passwords); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}


	UserID, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundedUser, err := queries.GetUserByID(UserID.UserID)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, passwords.OldPass)
	if !compareUserPassword {
		// Return, if password is not compare to stored in database.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "old password is not correct",
		})
	}

	hashedpass := utils.GeneratePassword(passwords.NewPass)

	user := &models.User{}
	db := database.Connection()
	err = db.First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "unable to process",
		})
	}

	user.PasswordHash = hashedpass
	db.Debug().Save(&user)

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "password changed",
	})

}



// @Tags User
// @Summary Create New User
// @Accept  json
// @Produce  json
// @Param data body models.SignUp true "User Data"
// @Success 201 {object} models.SignUpResponse
// @Failure 400
// @Router /user/signup [post]
func CreateUser(c *fiber.Ctx) error {

	newuser := &models.SignUp{}

	if err := c.BodyParser(newuser); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err := queries.AddUser(newuser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "user created",
	})
}
