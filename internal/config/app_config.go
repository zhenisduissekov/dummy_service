package config

import (
	"fmt"
	"github.com/ansrivas/fiberprometheus/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/zhenisduissekov/dummy_service/docs"
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
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/health", Health)
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
	fmt.Println("\n\n|", c.AuthApp.User, "|", c.AuthApp.Pass, "|\n\n")
	return app
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Error 401 {object} map[string]interface{}
// @Router /health [get]
func Health(f *fiber.Ctx) error {
	return f.Status(fiber.StatusOK).JSON(&fiber.Map{
		"code":    http.StatusOK,
		"message": "OK",
	})
}
