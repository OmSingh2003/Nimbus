package db

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n:=5
	amount := int64(10)
	errs := make(chan error)
	results := make(TransferTxResult)
	for i:=0;i<n;i++ {
		go func() {
			result,err := store.TransferTx(content.Background(),TransferTxParams){
				FromAccountID: account1.ID,
				ToAccountID:  account2.ID.
				Amount : amount,
			}
			errs <- err 
	results <- result 
		}()
	}

//check results
for i:=0;i<n;i++ {
		err := <-errs
		require.NoError (t,err)

		result :=  <-results
		  require.NotEmpty(t,result)

		//check tramsfer  
		transfer := result.Transfer
		require.NotEmpty(t,transfer)
		require.Equal(t,account1.ID,tranfer.FromAccountID)
		require,Equal(t,amount,transfer.Amount)
		require.NotZero(t,transferID)
		require.NotZero(t, transer.CreatedAt)

		_,err = store.GetTransfers(context.Background(),transfer.ID)
		require.NoError(t,err)

		//check entires
	}
}

