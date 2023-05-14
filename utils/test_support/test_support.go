package test_support

import (
	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/GandarfHSE/dentistryBackend/utils/migration"
)

func PrepareForTests() {
	config.LoadTestConfig()
	migration.MakeHardMigration()
}
