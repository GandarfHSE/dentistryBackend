package patient

type PatientInfo struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Passport string `json:"pass"`
}

type CreatePatientInfoRequest struct {
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Passport string `json:"pass"`
}

type CreatePatientInfoResponse struct {
}

type GetPatientInfoRequest struct {
	Uid int `json:"uid"`
}

type GetPatientInfoResponce struct {
	Info PatientInfo `json:"info"`
}
