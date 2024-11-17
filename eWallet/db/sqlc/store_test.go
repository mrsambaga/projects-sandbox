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
	numberOfTransactions := 5

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0;i < numberOfTransactions;i++ {
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

	for i:=0;i < numberOfTransactions;i++ {
		err := <-errs
		require.NoError(t, err)

		results := <-results
		require.NotNil(t, results)

		_, err = store.GetTransfer(context.Background(), results.Transfer.ID)
		require.NoError(t, err)

		transfer := results.Transfer
		require.NotNil(t, transfer)
		require.Equal(t, sender.ID, transfer.FromAccountID)
		require.Equal(t, receiver.ID, transfer.ToAccountID)
		require.NotZero(t, transfer.Amount)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetEntry(context.Background(), results.FromEntry.ID)
		require.NoError(t, err)

		fromEntry := results.FromEntry
		require.NotNil(t, fromEntry)
		require.Equal(t, sender.ID, fromEntry.AccountID)
		require.Equal(t, amount, fromEntry.Amount)

		_, err = store.GetEntry(context.Background(), results.ToEntry.ID)
		require.NoError(t, err)

		toEntry := results.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, receiver.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)

		receiverBalanceDiff := sender.Balance - results.FromAccount.Balance
		senderBalanceDiff := results.ToAccount.Balance - receiver.Balance
		require.Equal(t, receiverBalanceDiff, senderBalanceDiff)
		require.True(t, receiverBalanceDiff > 0)
	}
	updatedReceiver, err := testQueries.GetAccount(context.Background(), receiver.ID)
	require.NoError(t, err)

	updatedSemder, err := testQueries.GetAccount(context.Background(), sender.ID)
	require.NoError(t, err)

	require.Equal(t, receiver.Balance+int64(numberOfTransactions)*amount, updatedReceiver.Balance)
	require.Equal(t, sender.Balance-int64(numberOfTransactions)*amount, updatedSemder.Balance)
}