package db

import  (
	"testing"
	"context"
 "github.com/OmSingh2003/simple-bank/util"
	
	"github.com/stretchr/testify/require"
)
func TestCreateAccount(t *testing.T) {
	arg:= CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
}
account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
require.NotEmpty(t, account )

require.Equal(t,arg.Owner,account.Owner)
require.Equal(t, arg.Balance ,account.Balance)
require.Equal(t ,arg.Currency , account.Currency)

require.NotZero(t, account.ID)
	require.NotZero(t,account.CreatedAt)
}
