package main

import (
	"github.com/GandarfHSE/dentistryBackend/core"
	"github.com/rs/zerolog/log"
)

func main() {
	core.SetupLogs()

	log.Info().Msg("Getting daemon...")
	daemon := core.GetDaemon()
	log.Info().Msg("Registering handlers...")
	daemon.RegisterHandlers()
	log.Info().Msg("Summoning daemon...")
	daemon.SummonDaemon()
}
