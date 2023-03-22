package user

import (
	"github.com/GandarfHSE/dentistryBackend/core/auth"
	"github.com/GandarfHSE/dentistryBackend/utils/algo"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/ansel1/merry"
)

func CreateUserHandler(req CreateUserRequest, _ *cookie.Cookie) (*CreateUserResponse, merry.Error) {
	if doesUserExist(req.Login) {
		return nil, merry.New("User with this login already exists").WithHTTPCode(400)
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

func LoginHandler(req LoginRequest, _ *cookie.Cookie) (*LoginResponce, merry.Error) {
	if !doesUserExist(req.Login) {
		return nil, merry.New("User with this login does not exist").WithHTTPCode(400)
	}

	user := GetUserByLogin(req.Login)
	encodedPassword := algo.GenerateEncodedPassword(req.Password)
	if user.Password != encodedPassword {
		return nil, merry.New("Wrong login or password").WithHTTPCode(400)
	}

	authHandlers, err := auth.GetAuthHandlers()
	token, err := authHandlers.CreateToken(req.Login)
	if err != nil {
		return nil, merry.New("Cannot generate token!").WithHTTPCode(500)
	}

	return &LoginResponce{JWT: token}, nil
}
