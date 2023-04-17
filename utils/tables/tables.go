package tables

import (
	"fmt"
	"os"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/rs/zerolog/log"
)

var TableVersions = []string{
	createUserTable,
	createServiceTable,
}

func ApplyVersion(version int, s *database.Session) {
	err := database.Modify(s, TableVersions[version])
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("Error while applying version %d!", version))
		os.Exit(1)
	}
}
