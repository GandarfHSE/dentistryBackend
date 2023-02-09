package main

import "github.com/rs/zerolog/log"

func main() {
	core.setupLogs()
	log.Info().Msg("Hello!")
}
