package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuhengfdada/go-bank/util"
)

func getRandomUser() *User {
	return &User{
		Username:       util.GetRandomString(6),
		HashedPassword: util.GetRandomString(10),
		FullName:       util.GetRandomString(10),
		Email:          "hzshaoyuheng@gmail.com",
	}
}

func createRandomUser(t *testing.T) User {
	a := getRandomUser()
	usr, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Username:       a.Username,
		HashedPassword: a.HashedPassword,
		FullName:       a.FullName,
		Email:          a.Email,
	})
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, a.Username, usr.Username)
	require.Equal(t, a.HashedPassword, usr.HashedPassword)
	require.Equal(t, a.FullName, usr.FullName)
	require.Equal(t, a.Email, usr.Email)
	require.NotEmpty(t, usr.PasswordChangedAt)
	require.NotEmpty(t, usr.CreatedAt)
	return usr
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
