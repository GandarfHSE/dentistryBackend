package core

import (
	"fmt"
	"net/http"

	"github.com/GandarfHSE/dentistryBackend/utils/cli"
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
	var err error

	log.Info().Msg(fmt.Sprintf("Listening on %s:%s...", d.host, d.port))
	if cli.GetDoHTTPSFlag() {
		log.Info().Msg("Start a HTTPS server...")
		err = http.ListenAndServeTLS(url, config.GetHTTPSConfig().CertFile, config.GetHTTPSConfig().KeyFile, nil)
	} else {
		log.Info().Msg("Start a HTTP server...")
		err = http.ListenAndServe(url, nil)
	}

	if err != nil {
		log.Fatal().Err(err).Msg("Error in SummonDaemon")
	}
}
