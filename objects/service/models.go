package service

type Service struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Duration    int    `json:"duration"`
}

type CreateServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Duration    int    `json:"duration"`
}

// check README: empty json in response
type CreateServiceResponse struct {
	Err string `json:"err"`
}

type GetServiceListRequest struct {
}

type GetServiceListResponse struct {
	ServiceList []Service `json:"servicelist"`
}

type ServiceLink struct {
	Id  int `json:"id"`
	Did int `json:"did"`
	Sid int `json:"sid"`
}
