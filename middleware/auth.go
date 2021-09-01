package middleware

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/services"
	"strings"
)

func CheckToken(c *fiber.Ctx) error {
	if !strings.Contains(c.Path(), "articles") {
		return c.Next()
	}

	if c.Get("token") == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !services.IsTokenValid(c.Get("token")) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
