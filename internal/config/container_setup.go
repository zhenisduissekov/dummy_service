package config

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	abcService "github.com/zhenisduissekov/dummy_service/internal/abc/service"
	"github.com/zhenisduissekov/dummy_service/internal/db"
	defService "github.com/zhenisduissekov/dummy_service/internal/def/service"
	"github.com/zhenisduissekov/dummy_service/internal/tools"
)

type APIAuth struct {
	User string `mapstructure:"USER"`
	Pass string `mapstructure:"PASS"`
}

type AuthItems struct {
	Url  string `mapstructure:"URL"`
	User string `mapstructure:"USER"`
	Pass string `mapstructure:"PASS"`
}

type DBAuthItems struct {
	HostDB string `mapstructure:"HOST"`
	DB     string `mapstructure:"DB"`
	PortDB string `mapstructure:"PORT"`
	UserDB string `mapstructure:"USER"`
	PassDB string `mapstructure:"PASS"`
}

type Proxy struct {
	Address string `mapstructure:"Address"`
}

type Settings struct {
	AuthApp APIAuth        `mapstructure:"API_CREDS"`
	DBAuth  db.DBAuthItems `mapstructure:"DB_CREDS"`
	ABCAuth AuthItems      `mapstructure:"ABC"`
	DEFAuth AuthItems      `mapstructure:"DEF"`
	LVL     zerolog.Level  `mapstructure:"LOG_LEVEL"`
	Proxy   Proxy          `mapstructure:"PROXY"`
}

type Container struct {
	App        *fiber.App
	DBConn     *sqlx.DB
	ABCService *abcService.Service
	DEFService *defService.Service
	Proxy      *http.Client
}

func New() *Container {
	tools.PrepareLogging()

	conf, err := loadConfig(".")

	fmt.Println("cnfig", conf)

	if err != nil {
		log.Fatal().Err(err).Msg("could not load env")
	}

	zerolog.SetGlobalLevel(conf.LVL)

	db := conf.NewConnection()

	return &Container{
		App:        conf.NewApp(),
		ABCService: abcService.New(db, conf.ABCAuth.Url, conf.ABCAuth.User, conf.ABCAuth.Pass, conf.Proxy.Address),
		DEFService: defService.New(db, conf.DEFAuth.Url, conf.DEFAuth.User, conf.DEFAuth.Pass),
	}
}
