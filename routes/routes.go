package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"password-manager-backend/controllers"
	"password-manager-backend/errors"
	"password-manager-backend/middleware"
)

func Setup(app *fiber.App) {
	// TODO: encrypt passwords

	app.Use(cors.New(cors.Config{}))

	app.Use(errors.HandleException)
	app.Use(middleware.CheckToken)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome!")
	})

	api := app.Group("/api")

	user := api.Group("/user")
	user.Put("/register", controllers.Register)
	user.Post("/login", controllers.Login)
	user.Post("/logout", controllers.Logout)

	password := api.Group("/password", middleware.CheckToken)
	password.Get("/", controllers.GetPasswords)
	password.Get("/:name", controllers.GetPasswordsByName)
	password.Put("/", controllers.CreatePassword)
	password.Delete("/", controllers.DeletePassword)
	password.Patch("/", controllers.UpdatePassword)

	app.Use(middleware.LogOnMiddleWare)
}
