package controllers

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/errors"
	"password-manager-backend/services"
	"password-manager-backend/utils"
)

func Register(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "username", "password")
	if err != nil {
		return nil
	}

	c.JSON(services.CreateUser(data["username"], data["password"]))

	return c.Next()
}

func Login(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "username", "password")
	if err != nil {
		return nil
	}

	user := services.GetUser(data["username"])

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(errors.ErrorMessage{
			Error:   "BadRequest",
			Message: "User not found",
		})
	}

	if user.Password != data["password"] {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(errors.ErrorMessage{
			Error:   "BadRequest",
			Message: "Incorrect password",
		})
	}

	token := services.CreateToken(user.Id)

	c.JSON(fiber.Map{
		"token": token.Id,
	})

	return c.Next()
}

func Logout(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "tokenId")
	if err != nil {
		return nil
	}

	services.InvalidateToken(data["tokenId"])

	c.Status(fiber.StatusOK)
	c.JSON(fiber.Map{
		"status": "ok",
	})

	return c.Next()
}
