package core

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/rs/zerolog/log"
)

type Daemon struct {
	host string
	port string
}

func GetDaemon() *Daemon {
	serverConfig, err := config.GetServerConfig()
	if err != nil {
		log.Error().Err(err).Msg("Can't get server config!")
		os.Exit(1)
	}

	return &Daemon{
		host: serverConfig.Host,
		port: serverConfig.Port,
	}
}

func (d *Daemon) SummonDaemon() {
	url := d.host + ":" + d.port

	log.Info().Msg(fmt.Sprintf("Listening on %s:%s...", d.host, d.port))
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error in SummonDaemon")
		os.Exit(1)
	}
}
