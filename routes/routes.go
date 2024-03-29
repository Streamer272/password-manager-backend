package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"password-manager-backend/controllers"
	"password-manager-backend/errors"
	"password-manager-backend/logger"
	"password-manager-backend/middleware"
)

func Setup(app *fiber.App) {
	// TODO: encrypt passwords
	// TODO: add tests

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(middleware.LogOnMiddleWare)
	app.Use(errors.HandleException)
	app.Use(middleware.CheckToken)

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        50,
		Expiration: 0,
		LimitReached: func(c *fiber.Ctx) error {
			logger.LogError(fiber.ErrTooManyRequests)

			c.Status(fiber.StatusTooManyRequests)
			return c.JSON(errors.ErrorMessage{
				Error:   "TooManyRequests",
				Message: "",
			})
		},
	}))

	app.Get("/", controllers.Welcome)

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

	app.Use(func(c *fiber.Ctx) error {
		return nil
	})
}
