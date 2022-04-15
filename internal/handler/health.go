package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Health(f *fiber.Ctx) error {
	return f.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "OK",
	})
}
