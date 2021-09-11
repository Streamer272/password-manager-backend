package middleware

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/logger"
	"time"
)

func LogOnMiddleWare(c *fiber.Ctx) error {
	startTime := time.Now()

	dateTime := time.Now().Format("02-01-2006 15:04:05")

	err := c.Next()
	if err != nil {
		panic(err)
	}

	logger.Log(logger.BaseMessage, dateTime, "INFO", time.Since(startTime), c.Route().Method, c.Path())

	return nil
}
