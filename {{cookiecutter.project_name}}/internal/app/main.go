package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "{{cookiecutter.module_path}}/docs"
	"{{cookiecutter.module_path}}/internal/app/api"
)

// @title           {{cookiecutter.project_name}} API
// @version         1.0
// @description     Documentation for the {{cookiecutter.project_name}} API.

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					Token used to authenticate against the API.
func Run() error {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(helmet.New())

	// Prefork is only enabled in prod-mode.
	if !app.Config().Prefork {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	api.ConfigureRoutes(app)

	return app.Listen(":3000")
}
