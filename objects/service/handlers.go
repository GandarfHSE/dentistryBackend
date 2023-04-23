package service

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func CreateServiceHandler(req CreateServiceRequest, _ *cookie.Cookie) (*CreateServiceResponse, merry.Error) {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateServiceHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	exists, err := doesServiceExistByName(s, req.Name)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	if exists {
		return nil, merry.New("Service with this name already exists").WithHTTPCode(400)
	}

	err = addService(s, req)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateServiceResponse{}, nil
}

func GetServiceListHandler(req GetServiceListRequest, _ *cookie.Cookie) (*GetServiceListResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetServiceListHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	serviceList, err := getServiceList(s)
	if err != nil {
		log.Info().Msg(fmt.Sprintf(`%v`, serviceList))
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &GetServiceListResponse{ServiceList: serviceList}, nil
}
