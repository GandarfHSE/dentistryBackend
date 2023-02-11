package user

import (
	"github.com/ansel1/merry"
)

func CreateUserHandler(req CreateUserRequest) (*CreateUserResponse, error) {
	if doesUserExist(req.Login) {
		return nil, merry.New("User with this login already exist").WithHTTPCode(400)
	}
	if !IsRoleValid(req.Role) {
		return nil, merry.New("Invalid role").WithHTTPCode(400)
	}

	id, err := addUser(req)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateUserResponse{Id: id}, nil
}
