package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Server() {
	app := fiber.New()

	// logger
	app.Use(logger.New())

	// Setup routes
	SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
