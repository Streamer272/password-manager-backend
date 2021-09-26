package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	customLogger "password-manager-backend/logger"
	"password-manager-backend/models"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("password_manager_user:password_manager_password@tcp(localhost:3306)/password_manager"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	DB = conn

	err = DB.AutoMigrate(&models.User{}, &models.Password{}, &models.Token{})
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			customLogger.LogError(err)
		}
	}()
}
