package patient

type PatientInfo struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Passport string `json:"passport"`
}

type CreatePatientInfoRequest struct {
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Passport string `json:"passport"`
}

type CreatePatientInfoResponse struct {
}

type GetPatientInfoRequest struct {
	Uid int `json:"uid"`
}

type GetPatientInfoResponse struct {
	Info PatientInfo `json:"info"`
}
