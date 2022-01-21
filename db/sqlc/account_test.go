package db

import (
	"context"
	"learn/back-end/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)
}

func TestGetAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)
	gotAcc, err := testQueries.GetAccount(context.Background(), createdAcc.ID)
	require.NoError(t, err)

	require.Equal(t, createdAcc, gotAcc)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  1,
		Offset: 0,
		Owner:  "gqavfr",
	}
	accs, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accs, 1)

}

func TestUpdateAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAcc.ID,
		Balance: 10,
	}
	updatedAcc, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, updatedAcc.Balance, arg.Balance)
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account

}