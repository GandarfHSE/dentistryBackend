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

type CreateDoctorInfoResponce struct {
}

type GetDoctorInfoRequest struct {
	Uid int `json:"uid"`
}

type GetDoctorInfoResponce struct {
	Info DoctorInfo `json:"info"`
}
