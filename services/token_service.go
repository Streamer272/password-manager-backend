package services

import (
	"crypto/sha256"
	"encoding/hex"
	"password-manager-backend/database"
	"password-manager-backend/models"
	"strconv"
	"time"
)

func GetHashById(id uint) string {
	hash := sha256.New()
	hash.Write([]byte(strconv.Itoa(int(id))))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func CreateToken(userId uint) models.Token {
	deleteAllTokensByUserId(userId)

	// FIXME
	var highestToken models.Token
	database.Mysql.Model(&models.Token{}).Last(&highestToken)
	// till here

	token := models.Token{
		Uuid:    GetHashById(highestToken.Id + 1),
		Expires: time.Now().Add(time.Hour * 2).Unix(),
		UserId:  userId,
	}
	database.Mysql.Model(&models.Token{}).Create(&token)

	return token
}

func IsTokenValid(token interface{}) bool {
	var tokenModel models.Token
	database.Mysql.Model(&models.Token{}).Where("id = ?", token).First(&tokenModel)

	defer func() {
		if tokenModel.Expires <= time.Now().Unix() {
			database.Mysql.Model(&models.Token{}).Where("id = ?", token).Delete(&models.Token{})
		}
	}()

	if tokenModel.Id == 0 {
		return false
	}

	return tokenModel.Expires > time.Now().Unix()
}

func InvalidateToken(token interface{}) bool {
	var tokenModel models.Token
	database.Mysql.Model(&models.Token{}).Where("uuid = ?", token).First(&tokenModel)

	if tokenModel.Id == 0 {
		return false
	}

	database.Mysql.Model(&models.Token{}).Where("uuid = ?", token).Delete(&models.Token{})

	return true
}

func deleteAllTokensByUserId(userId uint) {
	database.Mysql.Model(&models.Token{}).Where("user_id = ?", userId).Delete(&models.Token{})
}
