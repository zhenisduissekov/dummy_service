package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s *Settings) NewConnection() *sqlx.DB {
	conn, err := s.DBAuth.Connect2DB()
	if err != nil {
		log.Fatal().Err(err).Msg("couldn't connect to DB")
	}
	return conn
}
