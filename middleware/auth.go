package middleware

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/services"
)

func CheckToken(c *fiber.Ctx) error {
	// FIXME

	if c.Get("tokenId") == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !services.IsTokenValid(c.Get("tokenId")) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
