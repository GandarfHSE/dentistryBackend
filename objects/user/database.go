package user

import (
	"github.com/GandarfHSE/dentistryBackend/utils/algo"
	"github.com/ansel1/merry"
)

// TODO: use postgre
var maxId int
var login_to_id = map[string]int{}
var id_to_user = map[int]User{}

func doesUserExist(login string) bool {
	id, exist := login_to_id[login]
	if !exist {
		return false
	}
	_, exist = id_to_user[id]
	return exist
}

// returns id of created user if succeeded
func addUser(req CreateUserRequest) (int, merry.Error) {
	login_to_id[req.Login] = maxId
	newUser := User{
		Id:       maxId,
		Login:    req.Login,
		Password: algo.GenerateEncodedPassword(req.Password),
		Role:     req.Role,
	}
	id_to_user[maxId] = newUser
	maxId++
	return maxId - 1, nil
}

func GetUserByLogin(login string) User {
	return id_to_user[login_to_id[login]]
}

func getUserList() ([]User, merry.Error) {
	arr := algo.Values(id_to_user)
	return arr, nil
}
