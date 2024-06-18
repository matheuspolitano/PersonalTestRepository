package db

import (
	"testing"

	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUse(t *testing.T) {
	firstName := utils.RandString(10)
	lastName := utils.RandString(10)
	email := utils.RandString(60)
	password := utils.RandString(10)

	user, err := testStore.CreateUser(firstName, lastName, email, password)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.FirstName, firstName)
	require.Equal(t, user.LastName, lastName)
	require.Equal(t, user.Email, email)

	isValid := utils.CheckPassword(password, user.HashPassword)
	require.True(t, isValid)

	err = testStore.DeleteById(user.UserID)
	require.NoError(t, err)
}
