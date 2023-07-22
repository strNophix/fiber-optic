package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// @Summary		Endpoint for testing latency
// @Description	Endpoint for testing latency
// @Tags 		Health
// @Success		200	{string} string
// @Router		/api/v1/ping [get]
func Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func ConfigureRoutes(router fiber.Router) error {
	v1 := router.Group("/v1")

	v1.Use(requestid.New())

	v1.Get("/ping", Ping)

	return nil
}
