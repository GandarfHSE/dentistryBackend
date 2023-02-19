package config

import (
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func GetServerConfig() (*ServerConfig, error) {
	if config == nil {
		log.Error().Msg("Reading ServerConfig from nil!")
		return nil, merry.New("Config hasn't been loaded!")
	}
	return &config.ServerConfig, nil
}

func GetAuthConfig() (*AuthConfig, error) {
	if config == nil {
		log.Error().Msg("Reading AuthConfig from nil!")
		return nil, merry.New("Config hasn't been loaded!")
	}
	return &config.AuthConfig, nil
}
