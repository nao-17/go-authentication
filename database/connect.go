package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("Cloud not connect the database")
	}
	DB = connection
	connection.AutoMigrate(&models.User{}, &models.PasswordReset{})
}
