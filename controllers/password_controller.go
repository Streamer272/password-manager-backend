package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/logger"
	"password-manager-backend/services"
)

func GetPasswords(c *fiber.Ctx) error {
	return c.JSON(services.GetAllPasswords(c.Get("tokenId")))
}

func GetPasswordsByName(c *fiber.Ctx) error {
	return c.JSON(services.GetPasswordsByName(c.Get("tokenId"), fmt.Sprintf("%v", c.Params("name"))))
}

func CreatePassword(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		logger.LogError(err)
		panic(fiber.ErrUnprocessableEntity)
	}

	return c.JSON(services.CreatePassword(c.Get("tokenId"), fmt.Sprintf("%v", data["name"]), fmt.Sprintf("%v", data["value"])))
}

func DeletePassword(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		logger.LogError(err)
		panic(fiber.ErrUnprocessableEntity)
	}

	return c.JSON(services.DeletePassword(c.Get("tokenId"), fmt.Sprintf("%v", data["name"])))
}
