package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
}

func TestCheckPassword_Missmatch(t *testing.T) {
	password := RandomString(6)
	wrongPassword := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)

	wrongHashedPassword, err := HashPassword(wrongPassword) 
	require.NoError(t, err)

	err = CheckPassword(wrongHashedPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestCheckPassword_Success(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password) 
	require.NoError(t, err)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}
