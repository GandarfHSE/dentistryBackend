package appointment

import "time"

type Appointment struct {
	Id        int       `json:"id"`
	Pid       int       `json:"pid"`
	Did       int       `json:"did"`
	Sid       int       `json:"sid"`
	Timebegin time.Time `json:"timebegin"`
	Timeend   time.Time `json:"timeend"`
}

type CreateAppointmentRequest struct {
	Pid  int       `json:"pid"`
	Did  int       `json:"did"`
	Sid  int       `json:"sid"`
	Time time.Time `json:"time"`
}

type CreateAppointmentPatientRequest struct {
	Did  int       `json:"did"`
	Sid  int       `json:"sid"`
	Time time.Time `json:"time"`
}

type CreateAppointmentDoctorRequest struct {
	Pid  int       `json:"pid"`
	Sid  int       `json:"sid"`
	Time time.Time `json:"time"`
}

type CreateAppointmentResponse struct {
}

type GetAppointmentByIdRequest struct {
	Id int `json:"id"`
}

type GetAppointmentByIdResponse struct {
	Appointment Appointment `json:"appointment"`
}
