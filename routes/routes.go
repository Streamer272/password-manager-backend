package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"password-manager-backend/controllers"
	"password-manager-backend/exceptions"
	"password-manager-backend/middleware"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{}))

	app.Use(exceptions.HandleException)
	app.Use(middleware.CheckToken)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome!")
	})

	api := app.Group("/api")
	user := api.Group("/user")

	user.Put("/register", controllers.Register)
	user.Post("/login", controllers.Login)
	user.Post("/logout", controllers.Logout)

	app.Use(middleware.LogOnMiddleWare)
}
