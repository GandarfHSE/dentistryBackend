package core

import (
	"net/http"

	"github.com/GandarfHSE/dentistryBackend/objects/doctor"
	"github.com/GandarfHSE/dentistryBackend/objects/patient"
	"github.com/GandarfHSE/dentistryBackend/objects/service"
	"github.com/GandarfHSE/dentistryBackend/objects/user"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello!\n"))
}

func RegisterHandlers() {
	http.HandleFunc("/hello", helloHandler)

	// user handlers
	createUserHandler := jsonHandlerWrapper(user.CreateUserHandler)
	http.HandleFunc("/user/create", createUserHandler)
	loginHandler := jsonHandlerWrapper(user.LoginHandler)
	http.HandleFunc("/user/login", loginHandler)
	getUserListHandler := noBodyHandlerWrapper(user.GetUserListHandler)
	http.HandleFunc("/user/list", getUserListHandler)
	whoAmIHandler := noBodyHandlerWrapper(user.WhoAmIHandler)
	http.HandleFunc("/user/whoami", whoAmIHandler)

	// patient wrappers
	createPatientInfoHandler := jsonHandlerWrapper(patient.CreatePatientInfoHandler)
	http.HandleFunc("/patient/create", createPatientInfoHandler)
	getPatientInfoHandler := jsonHandlerWrapper(patient.GetPatientInfoHandler)
	http.HandleFunc("/patient/get", getPatientInfoHandler)

	// doctor wrappers
	createDoctorInfoHandler := jsonHandlerWrapper(doctor.CreateDoctorInfoHandler)
	http.HandleFunc("/doctor/create", createDoctorInfoHandler)
	getDoctorInfoHandler := jsonHandlerWrapper(doctor.GetDoctorInfoHandler)
	http.HandleFunc("/doctor/get", getDoctorInfoHandler)
	findDoctorByNameSubstrHandler := jsonHandlerWrapper(doctor.FindDoctorByNameSubstrHandler)
	http.HandleFunc("/doctor/find/namesubstr", findDoctorByNameSubstrHandler)

	// service handlers
	createServiceHandler := jsonHandlerWrapper(service.CreateServiceHandler)
	http.HandleFunc("/service/create", createServiceHandler)
	getServiceListHandler := noBodyHandlerWrapper(service.GetServiceListHandler)
	http.HandleFunc("/service/list", getServiceListHandler)
}
