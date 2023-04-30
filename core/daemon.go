package core

import (
	"fmt"
	"net/http"

	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/rs/zerolog/log"
)

type Daemon struct {
	host string
	port string
}

func GetDaemon() *Daemon {
	serverConfig := config.GetServerConfig()

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
		log.Fatal().Err(err).Msg("Error in SummonDaemon")
	}
}
