package service

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/database"
)

func doesServiceExistByName(s *database.Session, name string) (bool, error) {
	_, err, exists := getServiceByName(s, name)
	return exists, err
}

func addService(s *database.Session, req CreateServiceRequest) error {
	q := `
		INSERT INTO "services" (name, description, cost, duration)
		VALUES ('%s', '%s', %d, %d);
	`

	return database.Modify(s, fmt.Sprintf(q, req.Name, req.Description, req.Cost, req.Duration))
}

// TODO: #12
func getServiceByName(s *database.Session, name string) (Service, error, bool) {
	q := `
		SELECT *
		FROM "services"
		WHERE "name" = '%s';
	`

	services, err := database.Get[Service](s, fmt.Sprintf(q, name))
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
		WHERE "id" = %d;
	`

	services, err := database.Get[Service](s, fmt.Sprintf(q, id))
	if err != nil {
		return Service{}, err, false
	}

	if len(services) > 0 {
		return services[0], nil, true
	} else {
		return Service{}, nil, false
	}
}

func getServiceList(s *database.Session) ([]Service, error) {
	q := `
		SELECT *
		FROM "services";
	`

	return database.Get[Service](s, q)
}
