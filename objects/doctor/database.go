package doctor

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
)

func addDoctorInfo(s *database.Session, req CreateDoctorInfoRequest) error {
	q := `
		INSERT INTO "doctors" (uid, name, post, exp)
		VALUES (%d, '%s', '%s', %d);
	`

	err := database.Modify(s, fmt.Sprintf(q, req.Uid, req.Name, req.Post, req.Exp))
	return err
}

func getDoctorInfo(s *database.Session, q string) (DoctorInfo, error, bool) {
	doctors, err := database.Get[DoctorInfo](s, q)
	if err != nil {
		return DoctorInfo{}, err, false
	}

	if len(doctors) > 0 {
		return doctors[0], nil, true
	} else {
		return DoctorInfo{}, nil, false
	}
}

func getDoctorInfoByUid(s *database.Session, uid int) (DoctorInfo, error, bool) {
	q := `
		SELECT *
		FROM "doctors"
		WHERE "uid" = '%v';
	`

	return getDoctorInfo(s, fmt.Sprintf(q, uid))
}

func findDoctorByNameSubstr(s *database.Session, name_substring string) ([]DoctorInfo, error) {
	q := `
		SELECT *
		FROM "doctors"
		WHERE "name" ~* '%s';
	`

	return database.Get[DoctorInfo](s, fmt.Sprintf(q, name_substring))
}
