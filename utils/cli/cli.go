package cli

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type CLIArgsModel struct {
	DoHardMigration bool
	DoHTTPS         bool
}

var _CLIArgs *CLIArgsModel

func ParseArgs() {
	_CLIArgs = &CLIArgsModel{}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "hard",
				Usage:       "Make hard database migration (drop all tables and create clean ones)",
				Value:       false,
				Required:    false,
				Destination: &_CLIArgs.DoHardMigration,
			},
			&cli.BoolFlag{
				Name:        "https",
				Usage:       "Run on https server instead of http",
				Value:       false,
				Required:    false,
				Destination: &_CLIArgs.DoHTTPS,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Can't parse command line args!")
	}
}

func GetDoHardMigrationFlag() bool {
	if _CLIArgs == nil {
		log.Fatal().Msg("Getting DoHardMigration flags, but flags have not been parsed!")
	}
	return _CLIArgs.DoHardMigration
}

func GetDoHTTPSFlag() bool {
	if _CLIArgs == nil {
		log.Fatal().Msg("Getting DoHTTPS flags, but flags have not been parsed!")
	}
	return _CLIArgs.DoHTTPS
}
