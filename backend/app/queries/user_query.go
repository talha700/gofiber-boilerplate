package queries

import (
	"errors"

	"github.com/gofiber-biolerplate/app/models"
	"github.com/gofiber-biolerplate/pkg/utils"
	"github.com/gofiber-biolerplate/platform/database"
	"gorm.io/gorm"
)

func GetUserByUserName(db *gorm.DB, data *models.SignIn) (models.User, error) {

	user := models.User{}

	db.Where("username = ?", data.Username).First(&user)

	var err error
	if (user == models.User{}) {
		err = errors.New("user not found")
	}
	return user, err
}

func GetUserByID(id string) (models.User, error) {

	db := database.Connection()
	user := models.User{}

	db.Where("id = ?", id).First(&user)

	var err error
	if (user == models.User{}) {
		err = errors.New("user not found")
	}
	return user, err
}

func AddUser(user *models.SignUp) error {

	db := database.Connection()

	hashedpass := utils.GeneratePassword(user.Password)

	err := db.Create(&models.User{Email: user.Email, Username: user.Username,
		UserRole: user.UserRole, PasswordHash: hashedpass}).Error

	if err != nil {
		err = errors.New("username or email already exists")
	}

	return err
}
