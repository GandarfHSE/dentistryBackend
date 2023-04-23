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

type CreateServiceResponse struct {
}

type GetServiceListRequest struct {
}

type GetServiceListResponse struct {
	ServiceList []Service `json:"servicelist"`
}
