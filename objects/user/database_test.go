package user_test

import (
	"testing"

	"github.com/GandarfHSE/dentistryBackend/objects/user"
	"github.com/GandarfHSE/dentistryBackend/utils/database"
	"github.com/GandarfHSE/dentistryBackend/utils/role"
	"github.com/GandarfHSE/dentistryBackend/utils/test_support"
	"github.com/stretchr/testify/require"
)

func TestAddGetUser(t *testing.T) {
	test_support.PrepareForTests()
	s, err := database.GetReadWriteSession()
	require.NoError(t, err)

	user_login := "test_login"

	_, err, exists := user.GetUserById(s, 1)
	require.NoError(t, err)
	require.False(t, exists)

	_, err, exists = user.GetUserByLogin(s, user_login)
	require.NoError(t, err)
	require.False(t, exists)

	add_req := user.CreateUserRequest{
		Login:    user_login,
		Password: "qwerty",
		Role:     1,
	}

	err = user.AddUser(s, add_req)
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
