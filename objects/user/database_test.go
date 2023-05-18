package user_test

import (
	"testing"

	"github.com/GandarfHSE/dentistryBackend/objects/user"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/GandarfHSE/dentistryBackend/utils/role"
	"github.com/GandarfHSE/dentistryBackend/utils/test_support"
	"github.com/stretchr/testify/require"
)

func createUser(s *database.Session, login string) error {
	add_req := user.CreateUserRequest{
		Login:    login,
		Password: "qwerty",
		Role:     1,
	}
	return user.AddUser(s, add_req)
}

func createAdmin(s *database.Session, login string) error {
	add_req := user.CreateUserRequest{
		Login:    login,
		Password: "12345",
		Role:     4,
		Keyword:  "kek",
	}
	return user.AddUser(s, add_req)
}

func TestAddGetUser(t *testing.T) {
	test_support.PrepareForTests()
	s, err := database.GetReadWriteSession()
	defer s.Close()
	require.NoError(t, err)

	user_login := "test_login"

	_, err, exists := user.GetUserById(s, 1)
	require.NoError(t, err)
	require.False(t, exists)

	_, err, exists = user.GetUserByLogin(s, user_login)
	require.NoError(t, err)
	require.False(t, exists)

	err = createUser(s, user_login)
	require.NoError(t, err)

	usr, err, exists := user.GetUserById(s, 1)
	require.NoError(t, err)
	require.True(t, exists)

	require.Equal(t, usr.Id, 1)
	require.Equal(t, usr.Login, user_login)
	require.Equal(t, usr.Role, role.Patient)

	usr, err, exists = user.GetUserByLogin(s, user_login)
	require.NoError(t, err)
	require.True(t, exists)

	require.Equal(t, usr.Id, 1)
	require.Equal(t, usr.Login, user_login)
	require.Equal(t, usr.Role, role.Patient)
}

func TestDoesUserExist(t *testing.T) {
	test_support.PrepareForTests()
	s, err := database.GetReadWriteSession()
	defer s.Close()
	require.NoError(t, err)

	login1 := "test_login"
	login2 := "login_test"

	exist, err := user.DoesUserExist(s, login1)
	require.NoError(t, err)
	require.False(t, exist)

	err = createUser(s, login1)
	require.NoError(t, err)

	exist, err = user.DoesUserExist(s, login1)
	require.NoError(t, err)
	require.True(t, exist)
	exist, err = user.DoesUserExist(s, login2)
	require.NoError(t, err)
	require.False(t, exist)
	exist, err = user.DoesUserWithUidExist(s, 1)
	require.NoError(t, err)
	require.True(t, exist)
	exist, err = user.DoesUserWithUidExist(s, 2)
	require.NoError(t, err)
	require.False(t, exist)

	err = createUser(s, login2)
	require.NoError(t, err)

	exist, err = user.DoesUserExist(s, login2)
	require.NoError(t, err)
	require.True(t, exist)
	exist, err = user.DoesUserWithUidExist(s, 2)
	require.NoError(t, err)
	require.True(t, exist)
}

func TestCheckUserRole(t *testing.T) {
	test_support.PrepareForTests()
	s, err := database.GetReadWriteSession()
	defer s.Close()
	require.NoError(t, err)

	err = createUser(s, "user_login")
	require.NoError(t, err)

	eq, err, ex := user.CheckUserRole(s, 1, role.Patient)
	require.NoError(t, err)
	require.True(t, ex)
	require.True(t, eq)

	eq, err, ex = user.CheckUserRole(s, 1, role.Doctor)
	require.NoError(t, err)
	require.True(t, ex)
	require.False(t, eq)

	eq, err, ex = user.CheckUserRole(s, 2, role.Admin)
	require.NoError(t, err)
	require.False(t, ex)

	err = createAdmin(s, "admin")
	require.NoError(t, err)

	eq, err, ex = user.CheckUserRole(s, 2, role.Admin)
	require.NoError(t, err)
	require.True(t, ex)
	require.True(t, eq)
}
