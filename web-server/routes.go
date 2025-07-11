package webserver

import (
	"encoding/json"
	"go-learn/data_structures"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type LinkedListWrapper struct {
	list data_structures.LinkedList
	mu   sync.RWMutex
}

var listWrapper = LinkedListWrapper{}

func ToJSON(data map[string]any) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func SetupRoutes(app *fiber.App) {
	// Main route:
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the LAAS aka (Linkedlist as a Service)!")
	})

	v1 := app.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		listWrapper.mu.RLock()
		defer listWrapper.mu.RUnlock()

		// Collect all nodes' data into a slice
		jsonData, err := listWrapper.list.Print_All()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error marshaling JSON: " + err.Error())
		}

		jsonBytes, ok := jsonData.([]byte)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("Error: failed to convert data to JSON bytes")
		}
		return c.Type("application/json").Send(jsonBytes)
	})

	v1.Post("/submit", func(c *fiber.Ctx) error {
		var data map[string]any
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON: " + err.Error())
		}

		listWrapper.mu.Lock()
		listWrapper.list.Append(data)
		listWrapper.mu.Unlock()

		return c.JSON(fiber.Map{"message": "Data received", "data": data})
	})

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
