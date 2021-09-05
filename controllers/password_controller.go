package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/services"
	"password-manager-backend/utils"
)

func GetPasswords(c *fiber.Ctx) error {
	c.JSON(services.GetAllPasswords(c.Get("tokenId")))

	return c.Next()
}

func GetPasswordsByName(c *fiber.Ctx) error {
	c.JSON(services.GetPasswordsByName(c.Get("tokenId"), fmt.Sprintf("%v", c.Params("name"))))

	return c.Next()
}

func CreatePassword(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "name", "value")
	if err != nil {
		return nil
	}

	c.JSON(services.CreatePassword(c.Get("tokenId"), fmt.Sprintf("%v", data["name"]), fmt.Sprintf("%v", data["value"])))

	return c.Next()
}

func DeletePassword(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "passwordId")
	if err != nil {
		return nil
	}

	c.JSON(services.DeletePassword(c.Get("tokenId"), fmt.Sprintf("%v", data["passwordId"])))

	return c.Next()
}

func UpdatePassword(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "passwordId", "name", "value")
	if err != nil {
		return nil
	}

	c.JSON(services.UpdatePassword(c.Get("tokenId"), data["passwordId"], fmt.Sprintf("%v", data["name"]), fmt.Sprintf("%v", data["value"])))

	return c.Next()
}
