package user

import (
	"errors"
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/algo"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/GandarfHSE/dentistryBackend/utils/role"
)

func AddUser(s *database.Session, req CreateUserRequest) error {
	q := `
		INSERT INTO "users" (login, password, role)
		VALUES ($1, $2, $3);
	`

	err := database.Modify(s, q, req.Login, algo.GenerateEncodedPassword(req.Password), req.Role)
	return err
}

func getUser(s *database.Session, q string) (User, error, bool) {
	users, err := database.Get[User](s, q)
	if err != nil {
		return User{}, err, false
	}

	if len(users) > 0 {
		return users[0], nil, true
	} else {
		return User{}, nil, false
	}
}

func GetUserByLogin(s *database.Session, login string) (User, error, bool) {
	q := `
		SELECT *
		FROM "users"
		WHERE "login" = '%s';
	`

	return getUser(s, fmt.Sprintf(q, login))
}

func GetUserById(s *database.Session, uid int) (User, error, bool) {
	q := `
		SELECT *
		FROM "users"
		WHERE "id" = '%v';
	`

	return getUser(s, fmt.Sprintf(q, uid))
}

func DoesUserExist(s *database.Session, login string) (bool, error) {
	_, err, exists := GetUserByLogin(s, login)
	return exists, err
}

func DoesUserWithUidExist(s *database.Session, uid int) (bool, error) {
	_, err, exists := GetUserById(s, uid)
	return exists, err
}

func getUserList(s *database.Session) ([]User, error) {
	q := `
		SELECT *
		FROM "users";
	`

	return database.Get[User](s, q)
}

func GetRoleFromCookie(s *database.Session, c *cookie.Cookie) (int, error) {
	user, err, exist := GetUserByLogin(s, c.Username)
	if err != nil {
		return role.Invalid, err
	}
	if !exist {
		return role.Invalid, errors.New(fmt.Sprintf("User with login %v does not exist", c.Username))
	}
	return user.Role, nil
}

func CheckUserRole(s *database.Session, uid int, role int) (bool, error) {
	user, err, exist := GetUserById(s, uid)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, errors.New(fmt.Sprintf("User with uid = %d does not exist!", uid))
	}
	return user.Role == role, nil
}
