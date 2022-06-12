package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuhengfdada/go-bank/util"
)

func getRandomAccount() *Account {
	return &Account{
		Owner:    util.GetRandomString(10),
		Balance:  util.GetRandomInt64(100, 1000),
		Currency: util.GetRandomString(3),
	}
}

func createRandomAccount(t *testing.T) Account {
	a := getRandomAccount()
	acc, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    a.Owner,
		Balance:  a.Balance,
		Currency: a.Currency,
	})
	if err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, acc.ID)
	require.Equal(t, a.Owner, acc.Owner)
	require.Equal(t, a.Balance, acc.Balance)
	require.Equal(t, a.Currency, acc.Currency)
	require.NotEmpty(t, acc.CreatedAt)
	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
