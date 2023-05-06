package patient

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
)

func addPatientInfo(s *database.Session, req CreatePatientInfoRequest) error {
	q := `
		INSERT INTO "patients" (uid, name, passport, photo)
		VALUES ($1, $2, $3, $4);
	`

	err := database.Modify(s, q, req.Uid, req.Name, req.Passport, req.Photo)
	return err
}

func getPatientInfo(s *database.Session, q string) (PatientInfo, error, bool) {
	patients, err := database.Get[PatientInfo](s, q)
	if err != nil {
		return PatientInfo{}, err, false
	}

	if len(patients) > 0 {
		return patients[0], nil, true
	} else {
		return PatientInfo{}, nil, false
	}
}

func getPatientInfoByUid(s *database.Session, uid int) (PatientInfo, error, bool) {
	q := `
		SELECT *
		FROM "patients"
		WHERE "uid" = %d;
	`

	return getPatientInfo(s, fmt.Sprintf(q, uid))
}
