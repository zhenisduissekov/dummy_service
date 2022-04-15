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

func main() {
	handler := handler.Handler{DI: DI}
	app := DI.App.Group("/api")
	{
		app.Post("/register", handler.Register)
		//api.Post("/schedulePayment", DIHandler.SchedulePayment) //todo: don't forget
	}

	log.Info().Msgf("Listening on %s", *listenAddress)
	if err := DI.App.Listen(*listenAddress); err != nil {
		log.Fatal().Err(err).Msg("Server down")
	}
}
