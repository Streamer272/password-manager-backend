package services

import (
	"password-manager-backend/database"
	"password-manager-backend/models"
	"time"
)

func CreateToken(userId uint) uint {
	token := models.Token{
		Expires: time.Now().Add(time.Hour * 2).Unix(),
		UserId:  userId,
	}
	database.DB.Model(&models.Token{}).Create(&token)

	return token.Id
}

func IsTokenValid(tokenId interface{}) bool {
	var token models.Token
	database.DB.Model(&models.Token{}).Where("id = ?", tokenId).First(&token)

	defer func() {
		if token.Expires <= time.Now().Unix() {
			database.DB.Model(&models.Token{}).Where("id = ?", tokenId).Delete(&models.Token{})
		}
	}()

	if token.Id == 0 {
		return false
	}

	return token.Expires > time.Now().Unix()
}

func InvalidateToken(tokenId interface{}) {
	database.DB.Model(&models.Token{}).Where("id = ?", tokenId).Delete(&models.Token{})
}
