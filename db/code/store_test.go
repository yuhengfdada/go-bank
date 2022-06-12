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
	for i := 0; i < 10; i++ {
		go func() {
			result, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Amount:        amount,
			})
			require.Nil(t, err)
			fmt.Printf("acc1: %v, acc2: %v\n", result.FromAccount.Balance, result.ToAccount.Balance)
			errchan <- err
		}()
	}
	for i := 0; i < 5; i++ {
		<-errchan
	}
	//require.Equal(t, acc1.Balance-amount, result.FromAccount.Balance)
	//require.Equal(t, acc2.Balance+amount, result.ToAccount.Balance)
}
