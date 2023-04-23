package patient

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/objects/user"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func CreatePatientInfoHandler(req CreatePatientInfoRequest, _ *cookie.Cookie) (*CreatePatientInfoResponse, merry.Error) {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreatePatientInfoHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	user_, err, exists := user.GetUserById(s, req.Uid)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !exists {
		return nil, merry.New(fmt.Sprintf("User with uid = %v does not exist!", req.Uid)).WithHTTPCode(400)
	}
	// TODO: #17
	if user_.Role != user.PatientRole {
		return nil, merry.New(fmt.Sprintf("User with uid = %v is not patient! Its role is %v", req.Uid, user_.Role)).WithHTTPCode(400)
	}

	err = addPatientInfo(s, req)
	return &CreatePatientInfoResponse{}, nil
}

func GetPatientInfoHandler(req GetPatientInfoRequest, _ *cookie.Cookie) (*GetPatientInfoResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetPatientInfoHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	doctorInfo, err, exists := getPatientInfoByUid(s, req.Uid)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !exists {
		return nil, merry.New(fmt.Sprintf("Patient info about uid = %v does not exist!", req.Uid)).WithHTTPCode(400)
	}

	return &GetPatientInfoResponse{Info: doctorInfo}, nil
}
