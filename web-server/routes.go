package webserver

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Main route:
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Web Server!")
	})

	v1 := app.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1.Post("/submit", func(c *fiber.Ctx) error {
		data := c.FormValue("data")
		return c.SendString("Data received: " + data)
	})

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
