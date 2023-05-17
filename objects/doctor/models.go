package doctor

type DoctorInfo struct {
	Id          int    `json:"id"`
	Uid         int    `json:"uid"`
	Name        string `json:"name"`
	Post        string `json:"post"`
	Exp         int    `json:"exp"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

type CreateDoctorInfoRequest struct {
	Uid         int    `json:"uid"`
	Name        string `json:"name"`
	Post        string `json:"post"`
	Exp         int    `json:"exp"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

// check README: empty json in response
type CreateDoctorInfoResponse struct {
	Err string `json:"err"`
}

type GetDoctorInfoRequest struct {
	Uid int `json:"uid"`
}

type GetDoctorInfoResponse struct {
	Info DoctorInfo `json:"info"`
}

type FindDoctorByNameSubstrRequest struct {
	Name string `json:"name"`
}

type FindDoctorByNameSubstrResponse struct {
	Result []DoctorInfo `json:"result"`
}
