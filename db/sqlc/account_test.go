package db

import (
	"context"
	"testing"

	"github.com/OmSingh2003/vaultguard-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	// Create a user first (required due to foreign key constraint)
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account 
	for i:=0;i<10;i++ {
		lastAccount = createRandomAccount(t)

	}
	arg := ListAccountsParams{
		Owner : lastAccount.Owner,
		Limit : 5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(),arg)
	require.NoError(t,err)
  require.NotEmpty(t,accounts)
	for _,account := range accounts {
		require.NotEmpty(t,account)
		require.Equal(t,lastAccount.Owner,account.Owner)
	}
}
