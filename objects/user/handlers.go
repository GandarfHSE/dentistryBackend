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

func GetUserListHandler(req GetUserListRequest, c *cookie.Cookie) (*GetUserListResponce, merry.Error) {
	if c == nil {
		return nil, merry.New("Unauthorized").WithHTTPCode(401)
	}

	requester := GetUserByLogin(c.Username)
	// TODO: maybe do separate logic with roles?
	if requester.Role != AdminRole && requester.Role != DevRole {
		return nil, merry.New("Access denied").WithHTTPCode(403)
	}

	userList, err := getUserList()
	if err != nil {
		return nil, merry.New("Can't get user list!").WithHTTPCode(500)
	}

	return &GetUserListResponce{UserList: userList}, nil
}
