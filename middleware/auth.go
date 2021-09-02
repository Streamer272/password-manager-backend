package middleware

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/errors"
	"password-manager-backend/services"
	"strings"
)

func CheckToken(c *fiber.Ctx) error {
	if !strings.Contains(c.Path(), "/api/password") {
		return c.Next()
	}

	if c.Get("tokenId") == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(errors.ErrorMessage{
			Error:   "Unauthorized",
			Message: "Missing `tokenId` header",
		})
	}

	if !services.IsTokenValid(c.Get("tokenId")) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(errors.ErrorMessage{
			Error:   "Unauthorized",
			Message: "Token is not valid",
		})
	}

	return c.Next()
}
