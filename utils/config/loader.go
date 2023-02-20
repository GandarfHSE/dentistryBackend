package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
)

var config *Config

func LoadConfig() error {
	configPath := os.Getenv("DENT_CONFIG_PATH")
	config = &Config{}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Error().Err(err).Msg("Can't open config file! Maybe DENT_CONFIG_PATH is unset...")
		return err
	}
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Error().Err(err).Msg("Can't read config file!")
		return err
	}

	err = json.Unmarshal(byteValue, config)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing json config!")
		return err
	}

	log.Info().Msg("Config has been read successfully!")
	return nil
}
