package doctor

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/objects/user"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func CreateDoctorInfoHandler(req CreateDoctorInfoRequest, _ *cookie.Cookie) (*CreateDoctorInfoResponse, merry.Error) {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateDoctorInfoHandler!")
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
	if user_.Role != user.DoctorRole {
		return nil, merry.New(fmt.Sprintf("User with uid = %v is not doctor! Its role is %v", req.Uid, user_.Role)).WithHTTPCode(400)
	}

	err = addDoctorInfo(s, req)
	return &CreateDoctorInfoResponse{}, nil
}

func GetDoctorInfoHandler(req GetDoctorInfoRequest, _ *cookie.Cookie) (*GetDoctorInfoResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetDoctorInfoHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	doctorInfo, err, exists := getDoctorInfoByUid(s, req.Uid)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !exists {
		return nil, merry.New(fmt.Sprintf("Doctor info about uid = %v does not exist!", req.Uid)).WithHTTPCode(400)
	}

	return &GetDoctorInfoResponse{Info: doctorInfo}, nil
}

func FindDoctorByNameSubstrHandler(req FindDoctorByNameSubstrRequest, _ *cookie.Cookie) (*FindDoctorByNameSubstrResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at FindDoctorByNameSubstrHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	doctors, err := findDoctorByNameSubstr(s, req.Name)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &FindDoctorByNameSubstrResponse{Result: doctors}, nil
}
