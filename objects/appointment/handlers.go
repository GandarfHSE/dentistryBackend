package appointment

import (
	"fmt"
	"time"

	"github.com/GandarfHSE/dentistryBackend/objects/service"
	"github.com/GandarfHSE/dentistryBackend/objects/user"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/GandarfHSE/dentistryBackend/utils/role"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func createAppointment(s *database.Session, pid int, did int, sid int, t time.Time) merry.Error {
	is_role_correct, err, exists := user.CheckUserRole(s, pid, role.Patient)
	if err != nil {
		return merry.Wrap(err).WithHTTPCode(500)
	}
	if !exists {
		return merry.New(fmt.Sprintf("User with uid = %d does not exist!", pid)).WithHTTPCode(400)
	}
	if !is_role_correct {
		return merry.New(fmt.Sprintf("User's role with uid = %d is not patient!", pid)).WithHTTPCode(400)
	}

	is_role_correct, err, exists = user.CheckUserRole(s, did, role.Doctor)
	if err != nil {
		return merry.Wrap(err).WithHTTPCode(500)
	}
	if !exists {
		return merry.New(fmt.Sprintf("User with uid = %d does not exist!", did)).WithHTTPCode(400)
	}
	if !is_role_correct {
		return merry.New(fmt.Sprintf("User's role with uid = %d is not doctor!", did)).WithHTTPCode(403)
	}

	exist, err := service.IsServiceExist(s, sid)
	if !exist {
		return merry.New(fmt.Sprintf("Service with sid = %d does not exist!", sid)).WithHTTPCode(403)
	}

	tend, err := service.GetServiceEndpoint(s, t, sid)
	if err != nil {
		return merry.Wrap(err).WithHTTPCode(500)
	}
	canCreate, err := canCreateAppointment(s, did, t, tend)
	if err != nil {
		return merry.Wrap(err).WithHTTPCode(500)
	}
	if !canCreate {
		return merry.New("Can't create an appointment!").WithHTTPCode(409)
	}

	err = createAppointmentDB(s, pid, did, sid, t, tend)
	if err != nil {
		return merry.Wrap(err).WithHTTPCode(500)
	}

	return nil
}

func CreateAppointmentHandler(req CreateAppointmentRequest, _ *cookie.Cookie) (*CreateAppointmentResponse, merry.Error) {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateAppointmentHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	merr := createAppointment(s, req.Pid, req.Did, req.Sid, req.Time)
	if merr != nil {
		return nil, merr
	}

	return &CreateAppointmentResponse{Err: "-"}, nil
}

func CreateAppointmentPatientHandler(req CreateAppointmentPatientRequest, c *cookie.Cookie) (*CreateAppointmentResponse, merry.Error) {
	if c == nil {
		return nil, merry.New("No cookie!").WithHTTPCode(401)
	}

	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateAppointmentHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	user_, err, exist := user.GetUserByLogin(s, c.Username)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !exist {
		return nil, merry.New(fmt.Sprintf("User with login %s does not exist!", c.Username)).WithHTTPCode(400)
	}

	merr := createAppointment(s, user_.Id, req.Did, req.Sid, req.Time)
	if merr != nil {
		return nil, merr
	}

	return &CreateAppointmentResponse{Err: "-"}, nil
}

func CreateAppointmentDoctorHandler(req CreateAppointmentDoctorRequest, c *cookie.Cookie) (*CreateAppointmentResponse, merry.Error) {
	if c == nil {
		return nil, merry.New("No cookie!").WithHTTPCode(401)
	}

	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateAppointmentHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	user_, err, exist := user.GetUserByLogin(s, c.Username)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !exist {
		return nil, merry.New(fmt.Sprintf("User with login %s does not exist!", c.Username)).WithHTTPCode(400)
	}

	merr := createAppointment(s, req.Pid, user_.Id, req.Sid, req.Time)
	if merr != nil {
		return nil, merr
	}

	return &CreateAppointmentResponse{Err: "-"}, nil
}

func GetAppointmentByIdHandler(req GetAppointmentByIdRequest, _ *cookie.Cookie) (*GetAppointmentByIdResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetAppointmentByIdHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	app, err, exists := getAppointmentById(s, req.Id)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !exists {
		return nil, merry.New(fmt.Sprintf("Appointment with id = %d does not exist!", req.Id)).WithHTTPCode(400)
	}
	return &GetAppointmentByIdResponse{Appointment: app}, nil
}

func getAppointmentList(q string) ([]Appointment, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at getAppointmentList!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	apps, err := database.Get[Appointment](s, q)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	return apps, nil
}

func GetAppointmentListHandler(req GetAppointmentListRequest, _ *cookie.Cookie) (*GetAppointmentListResponse, merry.Error) {
	q := `
		SELECT * FROM "appointments";
	`

	apps, err := getAppointmentList(q)
	if err != nil {
		return nil, err
	}
	return &GetAppointmentListResponse{AppointmentList: apps}, nil
}

func GetAppointmentListPatientHandler(req GetAppointmentListPatientRequest, _ *cookie.Cookie) (*GetAppointmentListResponse, merry.Error) {
	q := `
		SELECT * FROM "appointments"
		WHERE "pid" = %d;
	`

	apps, err := getAppointmentList(fmt.Sprintf(q, req.Pid))
	if err != nil {
		return nil, err
	}
	return &GetAppointmentListResponse{AppointmentList: apps}, nil
}

func GetAppointmentListDoctorHandler(req GetAppointmentListDoctorRequest, _ *cookie.Cookie) (*GetAppointmentListResponse, merry.Error) {
	q := `
		SELECT * FROM "appointments"
		WHERE "did" = %d;
	`

	apps, err := getAppointmentList(fmt.Sprintf(q, req.Did))
	if err != nil {
		return nil, err
	}
	return &GetAppointmentListResponse{AppointmentList: apps}, nil
}

func GetFreeTimeslotsHandler(req GetFreeTimeslotsRequest, _ *cookie.Cookie) (*GetFreeTimeslotsResponse, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetFreeTimeslotsHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	dateTime, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		log.Error().Err(err).Msg("Date parsing error in GetFreeTimeslotsHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	timeslots, err := getFreeTimeslots(s, req.Did, req.Sid, dateTime)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get free timeslots!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &GetFreeTimeslotsResponse{FreeTimeslots: timeslots}, nil
}
