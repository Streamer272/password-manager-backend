package middleware

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/logger"
	"password-manager-backend/services"
	"strings"
	"time"
)

func CheckToken(c *fiber.Ctx) error {
	if !strings.Contains(c.Path(), "/api/password") {
		return c.Next()
	}

	if c.Get("tokenId") == "" {
		logger.Log("token: %v", time.Now().Format("02-01-2006 15:04:05"), "INFO", "0", c.Get("tokenId"))
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !services.IsTokenValid(c.Get("tokenId")) {
		logger.Log("token: %v", time.Now().Format("02-01-2006 15:04:05"), "INFO", "0", c.Get("tokenId"))
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
