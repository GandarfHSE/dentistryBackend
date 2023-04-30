package migration

import (
	"fmt"
	"os"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/rs/zerolog/log"
)

type ver struct {
	Version int `json:"version"`
}

func createVersionTable(s *database.Session) {
	q := `
		CREATE TABLE "versions" (
			"version" INT NOT NULL
		);
	`

	err := database.Modify(s, q)
	if err != nil {
		log.Error().Err(err).Msg("Can't create version table!")
		os.Exit(1)
	}
}

func getVersion(s *database.Session) int {
	q := `
		SELECT *
		FROM "versions";
	`

	vs, err := database.Get[ver](s, q)
	if err != nil {
		log.Error().Err(err).Msg("Can't get database version in getVersion!")
		os.Exit(1)
	}
	if len(vs) != 1 {
		log.Error().Msg(fmt.Sprintf("getVersion: len(vs) == %v, but must equal to 1!", len(vs)))
		os.Exit(1)
	}

	return vs[0].Version
}

// remove old version (if exists) and set new
func setVersion(s *database.Session, v int) {
	q := `
		TRUNCATE TABLE "versions";
	`

	err := database.Modify(s, q)
	if err != nil {
		log.Error().Err(err).Msg("Can't get truncate version table!")
		os.Exit(1)
	}

	q = `
		INSERT INTO "versions" (version)
		VALUES ($1);
	`
	err = database.Modify(s, q, v)
	if err != nil {
		log.Error().Err(err).Msg("Can't set new version!")
		os.Exit(1)
	}
}
