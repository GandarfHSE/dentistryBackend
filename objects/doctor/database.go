package doctor

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
)

func addDoctorInfo(s *database.Session, req CreateDoctorInfoRequest) error {
	q := `
		INSERT INTO "doctors" (uid, name, post, exp, photo, description)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	err := database.Modify(s, q, req.Uid, req.Name, req.Post, req.Exp, req.Photo, req.Description)
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
		WHERE "uid" = %d;
	`

	return getDoctorInfo(s, fmt.Sprintf(q, uid))
}

func findDoctorByNameSubstr(s *database.Session, name_substring string) ([]DoctorInfo, error) {
	q := `
		SELECT *
		FROM "doctors"
		WHERE "name" ~* $1;
	`

	return database.Get[DoctorInfo](s, q, name_substring)
}
