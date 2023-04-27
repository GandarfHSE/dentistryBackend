package service

import (
	"time"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
)

func doesServiceExistByName(s *database.Session, name string) (bool, error) {
	_, err, exists := getServiceByName(s, name)
	return exists, err
}

func addService(s *database.Session, req CreateServiceRequest) error {
	q := `
		INSERT INTO "services" (name, description, cost, duration)
		VALUES ($1, $2, $3, $4);
	`

	return database.Modify(s, q, req.Name, req.Description, req.Cost, req.Duration)
}

// TODO: #12
func getServiceByName(s *database.Session, name string) (Service, error, bool) {
	q := `
		SELECT *
		FROM "services"
		WHERE "name" = $1;
	`

	services, err := database.Get[Service](s, q, name)
	if err != nil {
		return Service{}, err, false
	}

	if len(services) > 0 {
		return services[0], nil, true
	} else {
		return Service{}, nil, false
	}
}

func GetServiceById(s *database.Session, id int) (Service, error, bool) {
	q := `
		SELECT *
		FROM "services"
		WHERE "id" = $1;
	`

	services, err := database.Get[Service](s, q, id)
	if err != nil {
		return Service{}, err, false
	}

	if len(services) > 0 {
		return services[0], nil, true
	} else {
		return Service{}, nil, false
	}
}

func IsServiceExist(s *database.Session, id int) (bool, error) {
	_, err, exist := GetServiceById(s, id)
	return exist, err
}

func getServiceList(s *database.Session) ([]Service, error) {
	q := `
		SELECT *
		FROM "services";
	`

	return database.Get[Service](s, q)
}

func GetServiceEndpoint(s *database.Session, tbegin time.Time, sid int) (time.Time, error) {
	serv, err, exist := GetServiceById(s, sid)
	if err != nil {
		return time.Time{}, err
	}
	if !exist {
		return time.Time{}, merry.New("Service does not exist!")
	}

	return tbegin.Add(time.Minute * time.Duration(serv.Duration)), nil
}
