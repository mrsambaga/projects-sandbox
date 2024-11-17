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
	numberOfTransactions := 1

	errs := make(chan error)

	for i := 0;i < numberOfTransactions;i++ {
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: receiver.ID,
				ToAccountID: sender.ID,
				Ammount: amount,
			})

			errs <- err
		}()
	}

	for i:=0;i < numberOfTransactions;i++ {
		err := <-errs
		require.NoError(t, err)
	}
	updatedReceiver, err := testQueries.GetAccount(context.Background(), receiver.ID)
	require.NoError(t, err)

	updatedSemder, err := testQueries.GetAccount(context.Background(), sender.ID)
	require.NoError(t, err)

	require.Equal(t, receiver.Balance+int64(numberOfTransactions)*amount, updatedReceiver.Balance)
	require.Equal(t, sender.Balance-int64(numberOfTransactions)*amount, updatedSemder.Balance)
}