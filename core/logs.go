package core

import (
	"github.com/rs/zerolog"
)

func setupLogs() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
