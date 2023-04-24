package appointment

import "time"

type Appointment struct {
	Id   int       `json:"id"`
	Pid  int       `json:"pid"`
	Did  int       `json:"did"`
	Sid  int       `json:"sid"`
	Time time.Time `json:"time"`
}
