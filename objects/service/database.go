package service

import (
	"github.com/GandarfHSE/dentistryBackend/utils/algo"
	"github.com/ansel1/merry"
)

// TODO: use postgre
var maxId int
var name_to_id = map[string]int{}
var id_to_service = map[int]Service{}

func doesServiceExist(name string) bool {
	id, exist := name_to_id[name]
	if !exist {
		return false
	}
	_, exist = id_to_service[id]
	return exist
}

func addService(req CreateServiceRequest) (int, merry.Error) {
	name_to_id[req.Name] = maxId
	newService := Service{
		Id:          maxId,
		Name:        req.Name,
		Description: req.Description,
		Cost:        req.Cost,
		Duration:    req.Duration,
	}
	id_to_service[maxId] = newService
	maxId++
	return maxId - 1, nil
}

func getServiceList() ([]Service, merry.Error) {
	arr := algo.Values(id_to_service)
	return arr, nil
}
