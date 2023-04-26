package appointment

import (
	"time"

	"github.com/GandarfHSE/dentistryBackend/objects/service"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry/v2"
)

func timeFormat(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

func createAppointment(s *database.Session, pid int, did int, sid int, t1 time.Time, t2 time.Time) error {
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
		"timebegin" BETWEEN $2 AND $3 OR
		"timeend" BETWEEN $2 AND $3);
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

func getServiceEndpoint(s *database.Session, tbegin time.Time, sid int) (time.Time, error) {
	serv, err, exist := service.GetServiceById(s, sid)
	if err != nil {
		return time.Time{}, err
	}
	if !exist {
		return time.Time{}, merry.New("Service does not exist!")
	}

	return tbegin.Add(time.Minute * time.Duration(serv.Duration)), nil
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

func getAppointmentList(s *database.Session) ([]Appointment, error) {
	q := `
		SELECT * FROM "appointments";
	`

	apps, err := database.Get[Appointment](s, q)
	return apps, err
}
