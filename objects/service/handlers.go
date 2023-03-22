package service

import (
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/ansel1/merry"
)

func CreateServiceHandler(req CreateServiceRequest, _ *cookie.Cookie) (*CreateServiceResponce, merry.Error) {
	if doesServiceExist(req.Name) {
		return nil, merry.New("Service with this name already exists").WithHTTPCode(400)
	}

	id, err := addService(req)
	if err != nil {
		return nil, merry.New("Can't add new service!").WithHTTPCode(500)
	}

	return &CreateServiceResponce{Id: id}, nil
}