package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func ConfigureRoutes(router fiber.Router) error {
	v1 := router.Group("/v1")

	v1.Use(requestid.New())

	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	return nil
}
