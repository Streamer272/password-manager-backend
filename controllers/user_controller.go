package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/logger"
	"password-manager-backend/services"
	"strconv"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		logger.LogError(err)
		panic(fiber.ErrUnprocessableEntity)
	}

	return c.JSON(services.CreateUser(data["username"], data["password"]))
}

func Login(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		logger.LogError(err)
		panic(fiber.ErrUnprocessableEntity)
	}

	user := services.GetUser(data["username"])

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)

		return c.SendString("User not found")
	}

	if user.Password != data["password"] {
		c.Status(fiber.StatusBadRequest)

		return c.SendString("Incorrect password")
	}

	token := services.CreateToken(user.Id)

	return c.JSON(fiber.Map{
		"token": token.Id,
	})
}

func Logout(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		logger.LogError(err)
		panic(fiber.ErrUnprocessableEntity)
	}

	tokenId, _ := strconv.Atoi(fmt.Sprintf("%v", data["tokenId"]))

	services.InvalidateToken(uint(tokenId))

	return c.SendStatus(fiber.StatusOK)
}
