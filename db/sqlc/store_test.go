package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := testStore
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)
	 
	n := 5
	amount := int64(10)

	errs := make(chan error, n)
	results := make(chan TransferTXResult, n)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTX(context.Background(), TransferTXParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// Fetch transfer from DB
		fetchedTransfer, err := store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)
		require.Equal(t, transfer.ID, fetchedTransfer.ID)

		// Check fromEntry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)

		fetchedFromEntry, err := store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)
		require.Equal(t, fromEntry.ID, fetchedFromEntry.ID)

		// Check toEntry
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)

		fetchedToEntry, err := store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
		require.Equal(t, toEntry.ID, fetchedToEntry.ID)
		// check accounts 
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount 
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		//check account balances
	fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance 
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)
		k := int(diff1/amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}
	//check final updated balance 
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}

func TestTransferTxDeadlock(t *testing.T) {
	store := testStore
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">>>> Deadlock test - before:", account1.Balance, account2.Balance)
	
	n := 10
	amount := int64(10)

	errs := make(chan error, n)
	
	// Run n/2 transfers from account1 to account2
	// and n/2 transfers from account2 to account1 concurrently
	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID
		
		// Alternate direction every other transaction to create deadlock potential
		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		
		go func(fromID, toID int64) {
			_, err := store.TransferTX(context.Background(), TransferTXParams{
				FromAccountID: fromID,
				ToAccountID:   toID,
				Amount:        amount,
			})
			errs <- err
		}(fromAccountID, toAccountID)
	}
	
	// Collect results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}
	
	// Check final balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	
	fmt.Println(">>>> Deadlock test - after:", updatedAccount1.Balance, updatedAccount2.Balance)
	// Since equal transfers in both directions, balances should be unchanged
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}
