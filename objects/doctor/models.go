package doctor

type DoctorInfo struct {
	Id   int    `json:"id"`
	Uid  int    `json:"uid"`
	Name string `json:"name"`
	Post string `json:"post"`
	Exp  int    `json:"exp"`
}

type CreateDoctorInfoRequest struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
	Post string `json:"post"`
	Exp  int    `json:"exp"`
}

type CreateDoctorInfoResponse struct {
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
