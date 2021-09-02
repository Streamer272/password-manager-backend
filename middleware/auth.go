package middleware

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/services"
	"strings"
)

func CheckToken(c *fiber.Ctx) error {
	if !strings.Contains(c.Path(), "/api/password") {
		return c.Next()
	}

	if c.Get("tokenId") == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !services.IsTokenValid(c.Get("tokenId")) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
