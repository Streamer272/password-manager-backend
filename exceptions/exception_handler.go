package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/logger"
)

type ErrorJson struct {
	status int
	err    interface{}
	path   string
}

func HandleException(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			logger.LogError(err)

			var status int

			if err == fiber.ErrUnprocessableEntity {
				status = fiber.StatusUnprocessableEntity
			} else if err == fiber.ErrBadRequest {
				status = fiber.StatusBadRequest
			} else if err == fiber.ErrUnauthorized {
				status = fiber.StatusUnauthorized
			} else {
				status = fiber.StatusInternalServerError
			}

			c.SendStatus(status)
		}
	}()

	c.Next()

	return nil
}
