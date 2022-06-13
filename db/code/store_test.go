package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	amount := int64(10)
	fmt.Printf("acc1: %v, acc2: %v\n", acc1.Balance, acc2.Balance)
	errchan := make(chan error)
	reschan := make(chan TransferTxResult)
	for i := 0; i < 10; i++ {
		go func() {
			result, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Amount:        amount,
			})
			// fmt.Printf("acc1: %v, acc2: %v\n", result.FromAccount.Balance, result.ToAccount.Balance)
			errchan <- err
			reschan <- result
		}()
	}
	for i := 0; i < 5; i++ {
		err := <-errchan
		require.Nil(t, err)
		result := <-reschan
		require.Equal(t, (result.FromAccount.Balance-acc1.Balance)%amount, int64(0))
		require.Equal(t, (result.ToAccount.Balance-acc2.Balance)%amount, int64(0))
	}
}
