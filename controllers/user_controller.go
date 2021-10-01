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

	user, err := services.CreateUser(data["username"], data["password"])
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		err := c.JSON(errors.ErrorMessage{
			Error:   "BadRequest",
			Message: "User already exists",
		})
		if err != nil {
			panic(err)
		}
	} else {
		err = c.JSON(user)
		if err != nil {
			panic(err)
		}
	}

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

	err = c.JSON(fiber.Map{
		"token": token.Uuid,
	})
	if err != nil {
		panic(err)
	}

	return c.Next()
}

func Logout(c *fiber.Ctx) error {
	data, err := utils.CheckData(c, "token")
	if err != nil {
		return nil
	}

	succeeded := services.InvalidateToken(data["token"])

	if succeeded {
		c.Status(fiber.StatusOK)
		err = c.JSON(fiber.Map{
			"status": "ok",
		})
		if err != nil {
			panic(err)
		}
	} else {
		c.Status(fiber.StatusUnauthorized)
		err := c.JSON(errors.ErrorMessage{
			Error:   "Unauthorized",
			Message: "Token not valid",
		})
		if err != nil {
			panic(err)
		}
	}

	return c.Next()
}
