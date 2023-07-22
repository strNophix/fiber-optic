package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	v1 "{{cookiecutter.module_path}}/internal/app/api/v1"
)

func ConfigureRoutes(app *fiber.App) error {
	api := app.Group("/api")

	api.Use(helmet.New())

	err := v1.ConfigureRoutes(api)
	if err != nil {
		return err
	}

	return nil
}
