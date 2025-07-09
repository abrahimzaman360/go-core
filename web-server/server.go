package webserver

import "github.com/gofiber/fiber/v2"

func Server() {
	app := fiber.New()

	// Setup routes
	SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
