package database

import (
	"log"
	"os"

	"github.com/gofiber-biolerplate/app/models"
	"github.com/gofiber-biolerplate/pkg/utils"
	"gorm.io/gorm"
)

func CreateSuperUser(db *gorm.DB) {

	superuser := models.User{}
	user := os.Getenv("SUPER_USER")
	password := os.Getenv("SUPER_USER_PASSWORD")
	email := os.Getenv("SUPER_USER_EMAIL")

	db.Where("email = ?", email).First(&superuser)
	if superuser.Email == "" {
		log.Print("\nCreating User...")

		hashedpass := utils.GeneratePassword(password)

		db.Create(&models.User{Email: email, PasswordHash: hashedpass, Username: user,
			UserStatus: 1, UserRole: "admin"})

	} else {
		log.Print("Super User Already Exists")
	}
}
