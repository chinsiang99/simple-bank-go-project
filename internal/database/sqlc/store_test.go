package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	// Create two test accounts
	account1, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "alice",
		Balance:  1000,
		Currency: "USD",
	})
	require.NoError(t, err)

	account2, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "bob",
		Balance:  1000,
		Currency: "USD",
	})
	require.NoError(t, err)

	// Run n concurrent transfers: account1 -> account2
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	// Check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check transfer
		transfer := result.Transfer
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// Check entries
		fromEntry := result.FromEntry
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)

		toEntry := result.ToEntry
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)

		// Check accounts balances difference
		diff1 := account1.Balance - result.FromAccount.Balance
		diff2 := result.ToAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// Check final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTx_Deadlock(t *testing.T) {
	store := NewStore(testDB)

	// create 2 accounts
	accountA := createRandomAccount(t)
	accountB := createRandomAccount(t)

	// run concurrent transfers A->B and B->A
	n := 10
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := accountA.ID
		toAccountID := accountB.ID
		if i%2 == 1 {
			fromAccountID = accountB.ID
			toAccountID = accountA.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        10,
			})
			errs <- err
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		if err != nil {
			// Postgres will abort one transaction if deadlock occurs
			require.Contains(t, err.Error(), "deadlock detected")
		}
	}
}
