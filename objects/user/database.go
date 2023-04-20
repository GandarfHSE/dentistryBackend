package user

import (
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/algo"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
)

func addUser(s *database.Session, req CreateUserRequest) error {
	q := `
		INSERT INTO "users" (login, password, role)
		VALUES ('%s', '%s', %d);
	`

	err := database.Modify(s, fmt.Sprintf(q, req.Login, algo.GenerateEncodedPassword(req.Password), req.Role))
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

func getUserByLogin(s *database.Session, login string) (User, error, bool) {
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

func doesUserExist(s *database.Session, login string) (bool, error) {
	_, err, exists := getUserByLogin(s, login)
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
