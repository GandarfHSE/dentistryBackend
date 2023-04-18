package service

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func CreateServiceHandler(req CreateServiceRequest, _ *cookie.Cookie) (*CreateServiceResponce, merry.Error) {
	s, err := database.GetReadWriteSession()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateServiceHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	defer s.Close()

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

	return &CreateServiceResponce{}, nil
}

func GetServiceListHandler(req GetServiceListRequest, _ *cookie.Cookie) (*GetServiceListResponce, merry.Error) {
	s, err := database.GetReadSession()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetServiceListHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	defer s.Close()

	serviceList, err := getServiceList(s)
	if err != nil {
		log.Info().Msg(fmt.Sprintf(`%v`, serviceList))
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &GetServiceListResponce{ServiceList: serviceList}, nil
}
