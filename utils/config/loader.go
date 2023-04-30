package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
)

var config *Config

func LoadConfig() {
	configPath := os.Getenv("DENT_CONFIG_PATH")
	config = &Config{}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't open config file! Maybe DENT_CONFIG_PATH is unset...")
	}
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't read config file!")
	}

	err = json.Unmarshal(byteValue, config)
	if err != nil {
		log.Fatal().Err(err).Msg("Error occured while parsing json config!")
	}

	log.Info().Msg("Config has been read successfully!")
}
