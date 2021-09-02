package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/services"
	"strconv"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	return c.JSON(services.CreateUser(data["username"], data["password"]))
}

func Login(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
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
		"token": token,
	})
}

func Logout(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	tokenId, _ := strconv.Atoi(fmt.Sprintf("%v", data["tokenId"]))

	services.InvalidateToken(uint(tokenId))

	return c.SendStatus(fiber.StatusOK)
}
