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

type GetClinicListResponse struct {
	ClinicList []Clinic `json:"serviceList"`
}
