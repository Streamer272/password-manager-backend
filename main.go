package main

import (
	"github.com/gofiber/fiber/v2"
	"password-manager-backend/database"
	"password-manager-backend/routes"
)

func main() {
	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Password Manager",
	})

	app.Server().MaxConnsPerIP = 1

	routes.Setup(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
