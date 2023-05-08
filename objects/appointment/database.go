package appointment

import (
	"errors"
	"fmt"
	"time"

	"github.com/GandarfHSE/dentistryBackend/objects/service"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
)

func timeFormat(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

func createAppointmentDB(s *database.Session, pid int, did int, sid int, t1 time.Time, t2 time.Time) error {
	q := `
		INSERT INTO "appointments" (pid, did, sid, timebegin, timeend)
		VALUES ($1, $2, $3, $4, $5);
	`

	return database.Modify(s, q, pid, did, sid, t1, t2)
}

func getAppointmentsWithDoctorBetween(s *database.Session, did int, t1 time.Time, t2 time.Time) ([]Appointment, error) {
	q := `
		SELECT * FROM "appointments"
		WHERE "did" = $1 AND (
		("timebegin" >= $2 AND "timebegin" < $3) OR
		("timeend" > $2 AND "timeend" <= $3) OR
		("timebegin" <= $2 AND "timeend" >= $3));
	`

	apps, err := database.Get[Appointment](s, q, did, t1, t2)
	if err != nil || len(apps) == 0 {
		return nil, err
	}
	return apps, nil
}

func canCreateAppointment(s *database.Session, did int, t1 time.Time, t2 time.Time) (bool, error) {
	apps, err := getAppointmentsWithDoctorBetween(s, did, t1, t2)
	return apps == nil, err
}

func getAppointmentById(s *database.Session, id int) (Appointment, error, bool) {
	q := `
		SELECT * FROM "appointments"
		WHERE "id" = $1;
	`

	apps, err := database.Get[Appointment](s, q, id)
	if err != nil {
		return Appointment{}, err, false
	}

	if len(apps) > 0 {
		return apps[0], nil, true
	} else {
		return Appointment{}, nil, false
	}
}

func timeEqualOrLess(t1 time.Time, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}

func getFreeTimeslots(s *database.Session, did int, sid int, date time.Time) ([]Timeslot, error) {
	// [TODO: #33] Make custom working hours for doctors
	workingDayBegin := date.Add(time.Hour * time.Duration(9))
	workingDayEnd := date.Add(time.Hour * time.Duration(18))

	serv, err, exist := service.GetServiceById(s, sid)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New(fmt.Sprintf("Service with id %d does not exist!", sid))
	}

	var freeTimeslots []Timeslot
	serviceBegin := workingDayBegin
	for timeEqualOrLess(serviceBegin, workingDayEnd) {
		serviceEnd := serviceBegin.Add(time.Minute * time.Duration(serv.Duration))
		canCreate, err := canCreateAppointment(s, did, serviceBegin, serviceEnd)
		if err != nil {
			return nil, err
		}
		if timeEqualOrLess(serviceEnd, workingDayEnd) && canCreate {
			freeTimeslots = append(freeTimeslots, Timeslot{Timebegin: serviceBegin, Timeend: serviceEnd})
		}

		serviceBegin = serviceEnd
	}

	return freeTimeslots, nil
}
