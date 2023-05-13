package tables

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/rs/zerolog/log"
)

var TableVersions = []string{
	createUserTable,
	createServiceTable,
	createDoctorInfoTable,
	createPatientInfoTable,
	createAppointmentTable,
	createClinicTable,
	addDoctorInfoPhoto,
	addPatientInfoPhoto,
	createReviewTable,
	createServiceLinkTable,
}

func ApplyVersion(version int, s *database.Session) {
	err := database.Modify(s, TableVersions[version])
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Error while applying version %d!", version))
	}
}
