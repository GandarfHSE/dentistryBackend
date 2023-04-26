package appointment

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/objects/user"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

// TODO: remove copypaste
// TODO: check pid and did roles
func CreateAppointmentHandler(req CreateAppointmentRequest, _ *cookie.Cookie) (*CreateAppointmentResponse, merry.Error) {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateAppointmentHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	tend, err := getServiceEndpoint(s, req.Time, req.Sid)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	canCreate, err := canCreateAppointment(s, req.Did, req.Time, tend)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !canCreate {
		return nil, merry.New("Can't create an appointment!").WithHTTPCode(409)
	}

	err = createAppointment(s, req.Pid, req.Did, req.Sid, req.Time, tend)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateAppointmentResponse{}, nil
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
	// TODO: #17
	if user_.Role != user.PatientRole {
		return nil, merry.New("You are not a patient!").WithHTTPCode(403)
	}

	tend, err := getServiceEndpoint(s, req.Time, req.Sid)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	canCreate, err := canCreateAppointment(s, req.Did, req.Time, tend)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !canCreate {
		return nil, merry.New("Can't create an appointment!").WithHTTPCode(409)
	}

	err = createAppointment(s, user_.Id, req.Did, req.Sid, req.Time, tend)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateAppointmentResponse{}, nil
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
	// TODO: #17
	if user_.Role != user.DoctorRole {
		return nil, merry.New("You are not a doctor!").WithHTTPCode(403)
	}

	tend, err := getServiceEndpoint(s, req.Time, req.Sid)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	canCreate, err := canCreateAppointment(s, user_.Role, req.Time, tend)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !canCreate {
		return nil, merry.New("Can't create an appointment!").WithHTTPCode(409)
	}

	err = createAppointment(s, req.Pid, user_.Id, req.Sid, req.Time, tend)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateAppointmentResponse{}, nil
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
