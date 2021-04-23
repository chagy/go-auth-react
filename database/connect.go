package database

import (
	"go-auth-react/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/go_auth_react"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database!")
	}

	DB = connection

	connection.AutoMigrate(models.User{}, models.PasswordReset{})
}
