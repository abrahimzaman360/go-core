package webserver

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		data := c.FormValue("data")
		return c.SendString("Data received: " + data)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
