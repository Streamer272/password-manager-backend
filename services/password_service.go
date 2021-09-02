package services

import (
	"password-manager-backend/database"
	"password-manager-backend/models"
)

func GetAllPasswords(tokenId interface{}) []models.Password {
	var token models.Token
	database.DB.Model(&models.Token{}).Where("id = ?", tokenId).First(&token)

	var passwords []models.Password
	database.DB.Model(&models.Password{}).Where("user_id = ?", token.UserId).Find(&passwords)

	return passwords
}
