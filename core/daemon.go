package core

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type Daemon struct {
	host string
	port string
}

func GetDaemon() *Daemon {
	// TODO: use host/port from flags/envs
	return &Daemon{
		host: "localhost",
		port: "8083",
	}
}

func (d *Daemon) SummonDaemon() {
	url := d.host + ":" + d.port

	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error in SummonDaemon")
	}
}
