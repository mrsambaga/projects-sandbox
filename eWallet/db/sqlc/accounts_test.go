package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mrsambaga/projects-sandbox/eWallet/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	expectedResult := createRandomAccount(t)

	actualResult, err := testQueries.GetAccount(context.Background(), expectedResult.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actualResult)
	require.Equal(t, expectedResult.Owner, actualResult.Owner)
	require.Equal(t, expectedResult.Balance, actualResult.Balance)
	require.Equal(t, expectedResult.Currency, actualResult.Currency)
}

func TestUpdateAccount(t *testing.T) {
	existingAccount := createRandomAccount(t)
	updateArg := UpdateAccountParams{
		ID: existingAccount.ID,
		Balance: util.RandomMoney(),
	}

	actualResult, err := testQueries.UpdateAccount(context.Background(), updateArg)

	require.NoError(t, err)
	require.NotEmpty(t, actualResult)
	require.Equal(t, updateArg.Balance, actualResult.Balance)
	require.Equal(t, existingAccount.Owner, actualResult.Owner)
	require.Equal(t, existingAccount.Currency, actualResult.Currency)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	foundAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, foundAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	
	return account
}