package db

import (
	"context"
	"testing"

	"github.com/mrsambaga/projects-sandbox/eWallet/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	expectedResult := createRandomUser(t)

	actualResult, err := testQueries.GetUser(context.Background(), expectedResult.Username)

	require.NoError(t, err)
	require.NotEmpty(t, actualResult)
	require.Equal(t, expectedResult.Username, actualResult.Username)
	require.Equal(t, expectedResult.HashedPassword, actualResult.HashedPassword)
	require.Equal(t, expectedResult.FullName, actualResult.FullName)
	require.Equal(t, expectedResult.Email, actualResult.Email)
}

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomOwner(),
		HashedPassword: "secret",
		FullName: util.RandomOwner(),
		Email: util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user;
}
