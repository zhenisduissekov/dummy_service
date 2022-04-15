package config

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
			"code":    http.StatusOK,
			"message": "OK",
		})
	})

	app.Use(logger.New(logger.Config{
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${ip} ${url} ${bytesSent} ${bytesReceived} ${body}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 0,
		Output:       nil,
	}))

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			c.AuthApp.User: c.AuthApp.Pass,
		},
	}))

	return app
}
