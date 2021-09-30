package services

import (
	"password-manager-backend/database"
	"password-manager-backend/models"
)

func GetAllPasswords(token interface{}) []models.Password {
	var tokenModel models.Token
	database.Mysql.Model(&models.Token{}).Where("id = ?", token).First(&tokenModel)

	var passwords []models.Password
	database.Mysql.Model(&models.Password{}).Where("user_id = ?", tokenModel.UserId).Find(&passwords)

	return passwords
}

func GetPasswordsByName(token interface{}, name string) []models.Password {
	var tokenModel models.Token
	database.Mysql.Model(&models.Token{}).Where("id = ?", token).First(&tokenModel)

	var passwords []models.Password
	database.Mysql.Model(&models.Password{}).Where("user_id = ?", tokenModel.UserId).Where("name LIKE ?", "%"+name+"%").Find(&passwords)

	return passwords
}

func CreatePassword(token interface{}, name string, value string) models.Password {
	var tokenModel models.Token
	database.Mysql.Model(&models.Token{}).Where("id = ?", token).First(&tokenModel)

	password := models.Password{
		Name:   name,
		Value:  value,
		UserId: tokenModel.UserId,
	}
	database.Mysql.Model(&models.Password{}).Create(&password)

	return password
}

func DeletePassword(token interface{}, passwordId string) models.Password {
	var tokenModel models.Token
	database.Mysql.Model(&models.Token{}).Where("id = ?", token).First(&tokenModel)

	var password models.Password
	database.Mysql.Model(&models.Password{}).Where("user_id = ?", tokenModel.UserId).Where("id = ?", passwordId).First(&password)

	database.Mysql.Model(&models.Password{}).Where("user_id = ?", tokenModel.UserId).Where("id = ?", passwordId).Delete(&models.Password{})

	return password
}

func UpdatePassword(token interface{}, passwordId interface{}, name string, value string) models.Password {
	var tokenModel models.Token
	database.Mysql.Model(&models.Token{}).Where("id = ?", token).First(&tokenModel)

	// it doesn't work the other way
	database.Mysql.Model(&models.Password{}).Where("user_id = ?", tokenModel.UserId).Where("id = ?", passwordId).Update("name", name)
	database.Mysql.Model(&models.Password{}).Where("user_id = ?", tokenModel.UserId).Where("id = ?", passwordId).Update("value", value)

	var password models.Password
	database.Mysql.Model(&models.Password{}).Where("user_id = ?", tokenModel.UserId).Where("id = ?", passwordId).First(&password)

	return password
}
