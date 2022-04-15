package config

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	serviceName = "dummy_service"
)

func (c *Settings) NewApp() *fiber.App {
	app := fiber.New()
	prometheus := fiberprometheus.New(serviceName)
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
	app.Get("/health", func(f *fiber.Ctx) error {
		return f.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  http.StatusOK,
			"message": "OK",
		})
	})
	app.Use(prepareLogH, c.prepareAuthH)
	return app
}
