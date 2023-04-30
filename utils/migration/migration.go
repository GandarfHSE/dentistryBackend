package migration

import (
	"fmt"
	"os"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/GandarfHSE/dentistryBackend/utils/tables"
	"github.com/rs/zerolog/log"
)

func getAllTablesNames(s *database.Session) ([]string, error) {
	q := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema='public'
	`

	names, err := database.Get[string](s, q)
	if err != nil {
		return nil, err
	}

	return names, nil
}

func dropDatabase() error {
	s, err := database.GetReadWriteSession()
	if err != nil {
		return err
	}

	tables, err := getAllTablesNames(s)
	if err != nil {
		return err
	}

	for _, table_name := range tables {
		err = database.Modify(s, fmt.Sprintf("DROP TABLE \"%s\"", table_name))
		if err != nil {
			return err
		}
	}

	return s.Close()
}

func MakeFullMigration() {
	err := dropDatabase()
	if err != nil {
		log.Error().Err(err).Msg("Can't drop tables for full migration!")
		os.Exit(1)
	}

	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session for full migration!")
		os.Exit(1)
	}

	for version := range tables.TableVersions {
		tables.ApplyVersion(version, s)
	}
	createVersionTable(s)
	setVersion(s, len(tables.TableVersions)-1)
}

func MakeUpdateMigration() {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session for migrate migration!")
		os.Exit(1)
	}

	ver := getVersion(s)
	for version := range tables.TableVersions {
		if version > ver {
			tables.ApplyVersion(version, s)
		}
	}
	setVersion(s, len(tables.TableVersions)-1)
}
