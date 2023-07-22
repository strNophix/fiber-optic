package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"{{cookiecutter.module_path}}/internal/app/api"
)

func Run() error {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(compress.New())

	api.ConfigureRoutes(app)

	return app.Listen(":3000")
}
