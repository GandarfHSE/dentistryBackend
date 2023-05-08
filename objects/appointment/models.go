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

// check README: empty json in response
type CreateAppointmentResponse struct {
	Err string `json:"err"`
}

type GetAppointmentByIdRequest struct {
	Id int `json:"id"`
}

type GetAppointmentByIdResponse struct {
	Appointment Appointment `json:"appointment"`
}

type GetAppointmentListRequest struct {
}

type GetAppointmentListPatientRequest struct {
	Pid int `json:"pid"`
}

type GetAppointmentListDoctorRequest struct {
	Did int `json:"did"`
}

type GetAppointmentListResponse struct {
	AppointmentList []Appointment `json:"appointmentList"`
}

type Timeslot struct {
	Timebegin time.Time `json:"timebegin"`
	Timeend   time.Time `json:"timeend"`
}

type GetFreeTimeslotsRequest struct {
	Did  int    `json:"did"`
	Sid  int    `json:"sid"`
	Date string `json:"date"`
}

type GetFreeTimeslotsResponse struct {
	FreeTimeslots []Timeslot `json:"freeTimeslots"`
}
