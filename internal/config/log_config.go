package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func prepareLogH() fiber.Handler {
	return logger.New(logger.Config{
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${ip} ${url} ${bytesSent} ${bytesReceived} ${body}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 0,
		Output:       nil,
	})
}
