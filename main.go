package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/dummy_service/internal/config"
	"github.com/zhenisduissekov/dummy_service/internal/handler"
)

var (
	DI            *config.Container
	listenAddress = flag.String("port", ":1111", ":Ports")
)

func init() {
	DI = config.New()
}

// @title Fiber Swagger Example API
// @version 2.0
// @description A dummy service that takes a request, sends own request to another service and then conducts response back to user. Sample creds (test_user/test_pass)
// @termsOfService http://swagger.io/terms/

// @host localhost:1111
// @BasePath /
// @schemes http

// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func main() {
	handler := handler.Handler{DI: DI}

	app := DI.App.Group("/api")

	{
		app.Post("/register", handler.Register)
		app.Post("/schedulePayment", handler.Schedule)
	}

	log.Info().Msgf("Listening on %s", *listenAddress)
	if err := DI.App.Listen(*listenAddress); err != nil {
		log.Fatal().Err(err).Msg("Server down")
	}
}
