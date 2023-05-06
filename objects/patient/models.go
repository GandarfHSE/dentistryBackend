package patient

type PatientInfo struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Passport string `json:"passport"`
	Photo    string `json:"photo"`
}

type CreatePatientInfoRequest struct {
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Passport string `json:"passport"`
	Photo    string `json:"photo"`
}

// check README: empty json in response
type CreatePatientInfoResponse struct {
	Err string `json:"err"`
}

type GetPatientInfoRequest struct {
	Uid int `json:"uid"`
}

type GetPatientInfoResponse struct {
	Info PatientInfo `json:"info"`
}
