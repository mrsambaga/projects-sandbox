package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	sender := createRandomAccount(t)
	receiver := createRandomAccount(t)
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0;i < 5;i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: sender.ID,
				ToAccountID: receiver.ID,
				Ammount: amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i:=0;i<5;i++ {
		err := <-errs
		require.NoError(t, err)

		results := <-results
		require.NotNil(t, results)

		_, err = store.GetTransfer(context.Background(), results.Transfer.ID)
		require.NoError(t, err)

		_, err = store.GetEntry(context.Background(), results.FromEntry.ID)
		require.NoError(t, err)

		toEntry := results.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, receiver.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
	}
}