package clinic

type Clinic struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type CreateClinicRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

// check README: empty json in response
type CreateClinicResponse struct {
	Err string `json:"err"`
}

type GetClinicListRequest struct {
}

type ClinicListResponse struct {
	ClinicList []Clinic `json:"clinicList"`
}

type FindClinicByNameRequest struct {
	Name string `json:"name"`
}

type FindClinicByAddressRequest struct {
	Address string `json:"address"`
}

type FindClinicByPhoneRequest struct {
	Phone string `json:"phone"`
}

type FindClinicRequest struct {
	Str string `json:"str"`
}
