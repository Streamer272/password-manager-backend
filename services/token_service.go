package services

import (
	"password-manager-backend/database"
	"password-manager-backend/models"
	"time"
)

func CreateToken(userId uint) models.Token {
	token := models.Token{
		Expires: time.Now().Add(time.Hour * 2).Unix(),
		UserId:  userId,
	}
	database.DB.Model(&models.Token{}).Create(&token)

	return token
}

func IsTokenValid(token interface{}) bool {
	var tokenModel models.Token
	database.DB.Model(&models.Token{}).Where("id = ?", token).First(&tokenModel)

	defer func() {
		if tokenModel.Expires <= time.Now().Unix() {
			database.DB.Model(&models.Token{}).Where("id = ?", token).Delete(&models.Token{})
		}
	}()

	if tokenModel.Id == 0 {
		return false
	}

	return tokenModel.Expires > time.Now().Unix()
}

func InvalidateToken(token interface{}) {
	database.DB.Model(&models.Token{}).Where("id = ?", token).Delete(&models.Token{})
}
