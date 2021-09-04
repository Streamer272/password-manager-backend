package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/errors"
	"password-manager-backend/logger"
)

func CheckData(c *fiber.Ctx, requiredData ...string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		logger.LogError(err)
		fmt.Printf("err: %v\n", err)

		c.Status(fiber.StatusUnprocessableEntity)
		c.JSON(errors.ErrorMessage{
			Error:   "UnprocessableEntity",
			Message: "Couldn't parse JSON",
		})

		return data, fiber.ErrUnprocessableEntity
	}

	for _, required := range requiredData {
		if data[required] == nil {
			c.Status(fiber.StatusUnprocessableEntity)
			c.JSON(errors.ErrorMessage{
				Error:   "UnprocessableEntity",
				Message: "Missing " + required,
			})
			fmt.Printf("missing %v (%v)\n", required, data[required])

			return data, fiber.ErrUnprocessableEntity
		}
	}

	return data, nil
}
