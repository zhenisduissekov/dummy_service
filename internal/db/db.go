package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/rs/zerolog/log"
)

var (
	dbConnString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

type DBAuthItems struct {
	HostDB     string `mapstructure:"HOST"`
	DB         string `mapstructure:"DB"`
	PortDB     string `mapstructure:"PORT"`
	UsernameDB string `mapstructure:"USER"`
	PasswordDB string `mapstructure:"PASS"`
	DriverName string `mapstructure:"DRIVER"`
}

func (db *DBAuthItems) Connect2DB() (*sqlx.DB, error) {
	connStr := fmt.Sprintf(dbConnString, db.HostDB, db.PortDB, db.UsernameDB, db.PasswordDB, db.DB)
	conn, err := sqlx.Connect(db.DriverName, connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("error at sqlx.Open")
		return nil, err
	}
	log.Trace().Msg("connected to DB")

	if err := goose.Up(conn.DB, "migrations"); err != nil {
		log.Fatal().Err(err).Msg("error at goose.Up")
		panic(err)
	}

	return conn, err
}
