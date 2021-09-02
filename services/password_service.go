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

func GetPasswordsByName(tokenId interface{}, name string) []models.Password {
	var token models.Token
	database.DB.Model(&models.Token{}).Where("id = ?", tokenId).First(&token)

	var passwords []models.Password
	database.DB.Model(&models.Password{}).Where("user_id = ?", token.UserId).Where("name LIKE ?", "%"+name+"%").Find(&passwords)

	return passwords
}

func CreatePassword(tokenId interface{}, name string, value string) models.Password {
	var token models.Token
	database.DB.Model(&models.Token{}).Where("id = ?", tokenId).First(&token)

	password := models.Password{
		Name:   name,
		Value:  value,
		UserId: token.UserId,
	}
	database.DB.Model(&models.Password{}).Create(&password)

	return password
}

func DeletePassword(tokenId interface{}, name string) models.Password {
	var token models.Token
	database.DB.Model(&models.Token{}).Where("id = ?", tokenId).First(&token)

	var password models.Password
	database.DB.Model(&models.Password{}).Where("user_id = ?", token.UserId).Where("name LIKE ?", "%"+name+"%").First(&password)

	database.DB.Model(&models.Password{}).Where("user_id = ?", token.UserId).Where("name LIKE ?", "%"+name+"%").Delete(&models.Password{})

	return password
}
