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

type CreateServiceResponce struct {
}

type GetServiceListRequest struct {
}

type GetServiceListResponce struct {
	ServiceList []Service `json:"servicelist"`
}
