package test_support

import (
	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/GandarfHSE/dentistryBackend/utils/migration"
	"github.com/rs/zerolog"
)

func PrepareForTests() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel) // temporary WA, see #42
	config.LoadTestConfig()
	migration.MakeHardMigration()
}
