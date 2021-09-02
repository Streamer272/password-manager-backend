package services

import (
	"fmt"
	"password-manager-backend/database"
	"password-manager-backend/models"
)

func CreateUser(username interface{}, password interface{}) models.User {
	user := models.User{
		Username: fmt.Sprintf("%v", username),
		Password: fmt.Sprintf("%v", password),
	}
	database.DB.Model(&models.User{}).Create(&user)

	return user
}

func GetUser(username interface{}) models.User {
	var user models.User
	database.DB.Model(&models.User{}).Where("username = ?", username).First(&user)

	return user
}
