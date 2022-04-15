package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func loadConfig(path string) (*Settings, error) {
	var CNT *Settings
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Err(err).Msg("error at ReadConfig")
		return CNT, err
	}
	err = viper.Unmarshal(&CNT)
	if err != nil {
		log.Err(err).Msg("error at config")
	}

	log.Trace().Msgf("Configuration loaded: %v", CNT)
	return CNT, err
}
