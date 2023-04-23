package user

import (
	"github.com/GandarfHSE/dentistryBackend/core/auth"
	"github.com/GandarfHSE/dentistryBackend/utils/algo"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

func CreateUserHandler(req CreateUserRequest, _ *cookie.Cookie) (*CreateUserResponse, merry.Error) {
	s, err := database.GetReadWriteSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get write session at CreateUserHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	exists, err := doesUserExist(s, req.Login)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if exists {
		return nil, merry.New("User with this login already exists").WithHTTPCode(400)
	}
	if !IsRoleValid(req.Role) {
		return nil, merry.New("Invalid role").WithHTTPCode(400)
	}

	if err = addUser(s, req); err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &CreateUserResponse{}, nil
}

func LoginHandler(req LoginRequest, _ *cookie.Cookie) (*LoginResponce, merry.Error) {
	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at LoginHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	user, err, exists := getUserByLogin(s, req.Login)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	if !exists {
		return nil, merry.New("User with this login does not exist").WithHTTPCode(400)
	}

	encodedPassword := algo.GenerateEncodedPassword(req.Password)
	if user.Password != encodedPassword {
		return nil, merry.New("Wrong login or password").WithHTTPCode(400)
	}

	authHandlers, err := auth.GetAuthHandlers()
	token, err := authHandlers.CreateToken(req.Login)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &LoginResponce{JWT: token}, nil
}

func GetUserListHandler(req GetUserListRequest, c *cookie.Cookie) (*GetUserListResponce, merry.Error) {
	if c == nil {
		return nil, merry.New("No cookie!").WithHTTPCode(401)
	}

	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetUserListHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	requester, err, _ := getUserByLogin(s, c.Username)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}
	// TODO: #17
	if requester.Role != AdminRole && requester.Role != DevRole {
		return nil, merry.New("Access denied").WithHTTPCode(403)
	}

	userList, err := getUserList(s)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &GetUserListResponce{UserList: userList}, nil
}

func WhoAmIHandler(req WhoAmIRequest, c *cookie.Cookie) (*WhoAmIResponse, merry.Error) {
	if c == nil {
		return nil, merry.New("No cookie!").WithHTTPCode(401)
	}

	s, err := database.GetReadSession()
	defer s.Close()
	if err != nil {
		log.Error().Err(err).Msg("Can't get read session at GetUserListHandler!")
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	requester, err, _ := getUserByLogin(s, c.Username)
	if err != nil {
		return nil, merry.Wrap(err).WithHTTPCode(500)
	}

	return &WhoAmIResponse{Id: requester.Id, Login: requester.Login, Role: requester.Role}, nil
}
