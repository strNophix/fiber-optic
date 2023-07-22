package v1

import (
	"github.com/gofiber/fiber/v2"
)

func ConfigureRoutes(router fiber.Router) error {
	v1 := router.Group("/v1")

	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	return nil
}
