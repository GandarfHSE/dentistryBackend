package config

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func tryLoadConfig() error {
	log.Info().Msg("Trying to load config...")
	err := LoadConfig()
	if err != nil {
		log.Error().Msg("Failed to load config!")
		return err
	}
	return nil
}

func GetServerConfig() (*ServerConfig, error) {
	if config == nil {
		err := tryLoadConfig()
		if err != nil {
			return nil, err
		}
	}
	return &config.ServerConfig, nil
}

func GetAuthConfig() (*AuthConfig, error) {
	if config == nil {
		err := tryLoadConfig()
		if err != nil {
			return nil, err
		}
	}
	return &config.AuthConfig, nil
}

func GetDatabaseConfig() (*DBConfig, error) {
	if config == nil {
		err := tryLoadConfig()
		if err != nil {
			return nil, err
		}
	}
	return &config.DatabaseConfig, nil
}

// TODO - move it
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
