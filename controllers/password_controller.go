package controllers

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/services"
)

func GetPasswords(c *fiber.Ctx) error {
	return c.JSON(services.GetAllPasswords(c.Get("tokenId")))
}
