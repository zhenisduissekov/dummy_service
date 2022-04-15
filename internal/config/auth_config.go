package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func (s *Settings) prepareAuthH() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			s.AuthApp.user: s.AuthApp.pass,
		},
		Realm: "Forbidden",
	})
}
