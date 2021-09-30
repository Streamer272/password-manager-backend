package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/database"
	"password-manager-backend/models"
)

func CreateUser(username interface{}, password interface{}) (models.User, error) {
	user := models.User{
		Username: fmt.Sprintf("%v", username),
		Password: fmt.Sprintf("%v", password),
	}
	database.Mysql.Model(&models.User{}).Create(&user)

	if user.Id == 0 {
		return user, fiber.ErrBadRequest
	}

	return user, nil
}

func GetUser(username interface{}) models.User {
	var user models.User
	database.Mysql.Model(&models.User{}).Where("username = ?", username).First(&user)

	return user
}
