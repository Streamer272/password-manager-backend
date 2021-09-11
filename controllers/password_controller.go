package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/services"
	"password-manager-backend/utils"
)

func GetPasswords(c *fiber.Ctx) error {
	err := c.JSON(services.GetAllPasswords(c.Get("token")))
	if err != nil {
		panic(err)
	}

	return c.Next()
}

func GetPasswordsByName(c *fiber.Ctx) error {
	err := c.JSON(services.GetPasswordsByName(c.Get("token"), fmt.Sprintf("%v", c.Params("name"))))
	if err != nil {
		panic(err)
	}

	return c.Next()
}

func CreatePassword(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "name", "value")
	if err != nil {
		return nil
	}

	err = c.JSON(services.CreatePassword(c.Get("token"), fmt.Sprintf("%v", data["name"]), fmt.Sprintf("%v", data["value"])))
	if err != nil {
		panic(err)
	}

	return c.Next()
}

func DeletePassword(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "passwordId")
	if err != nil {
		return nil
	}

	err = c.JSON(services.DeletePassword(c.Get("token"), fmt.Sprintf("%v", data["passwordId"])))
	if err != nil {
		panic(err)
	}

	return c.Next()
}

func UpdatePassword(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "passwordId", "name", "value")
	if err != nil {
		return nil
	}

	err = c.JSON(services.UpdatePassword(c.Get("token"), data["passwordId"], fmt.Sprintf("%v", data["name"]), fmt.Sprintf("%v", data["value"])))
	if err != nil {
		panic(err)
	}

	return c.Next()
}
