package clinic

import (
	"github.com/GandarfHSE/dentistryBackend/objects/user"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/GandarfHSE/dentistryBackend/utils/role"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func CreateClinicHandler(req CreateClinicRequest, c *cookie.Cookie) (*CreateClinicResponse, merry.Error) {
	if c == nil {
		return nil, merry.New("No cookie!").WithHTTPCode(401)
	}

	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateClinicHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	user_role, err := user.GetRoleFromCookie(s, c)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !role.IsRoleAtLeast(user_role, role.Admin) {
		return nil, merry.New("Clinic creation allowed only for admins!").WithHTTPCode(403)
	}

	err = createClinic(s, req)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateClinicResponse{Err: "-"}, nil
}

func GetClinicListHandler(req GetClinicListRequest, _ *cookie.Cookie) (*GetClinicListResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetClinicListHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	q := `
		SELECT * FROM "clinics";
	`
	clinics, err := database.Get[Clinic](s, q)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &GetClinicListResponse{ClinicList: clinics}, nil
}
