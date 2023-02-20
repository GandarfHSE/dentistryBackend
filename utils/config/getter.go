package config

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func GetServerConfig() (*ServerConfig, error) {
	if config == nil {
		log.Warn().Msg("Reading ServerConfig from nil: trying to load config...")
		err := LoadConfig()
		if err != nil {
			log.Error().Msg("Reading ServerConfig from nil: failed to load config!")
			return nil, err
		}
	}
	return &config.ServerConfig, nil
}

func GetAuthConfig() (*AuthConfig, error) {
	if config == nil {
		log.Warn().Msg("Reading AuthConfig from nil: trying to load config...")
		err := LoadConfig()
		if err != nil {
			log.Error().Msg("Reading AuthConfig from nil: failed to load config!")
			return nil, err
		}
	}
	return &config.AuthConfig, nil
}

func GetAbsPrivatePath() (string, error) {
	authConfig, err := GetAuthConfig()
	if err != nil {
		return "", err
	}

	absolutePrivatePath, err := filepath.Abs(*&authConfig.PrivatePath)
	if err != nil {
		return "", err
	}

	return absolutePrivatePath, nil
}

func GetAbsPublicPath() (string, error) {
	authConfig, err := GetAuthConfig()
	if err != nil {
		return "", err
	}

	absolutePublicPath, err := filepath.Abs(*&authConfig.PublicPath)
	if err != nil {
		return "", err
	}

	return absolutePublicPath, nil
}
