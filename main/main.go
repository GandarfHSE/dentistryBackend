package main

import (
	"github.com/GandarfHSE/dentistryBackend/core"
	"github.com/GandarfHSE/dentistryBackend/core/auth"

	"github.com/GandarfHSE/dentistryBackend/utils/cli"
	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/GandarfHSE/dentistryBackend/utils/migration"
	"github.com/rs/zerolog/log"
)

func main() {
	core.SetupLogs()

	log.Info().Msg("Parsing command line flags...")
	cli.ParseArgs()

	log.Info().Msg("Loading config...")
	config.LoadConfig()

	log.Info().Msg("Loading authHandlers...")
	auth.LoadAuthHandlers()

	log.Info().Msg("Registering handlers...")
	core.RegisterHandlers()

	log.Info().Msg("Getting daemon...")
	daemon := core.GetDaemon()

	log.Info().Msg("Making migration...")
	migration.MakeMigration()

	log.Info().Msg("Summoning daemon...")
	daemon.SummonDaemon()
}
