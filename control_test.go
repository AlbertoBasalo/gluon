package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCalculateBalance(t *testing.T) {
	t.Run("it should calculate the balance", func(t *testing.T) {
		// arrange
		var transactions []Transaction
		var aDeposit Transaction = Transaction{
			date:            time.Now(),
			amount:          0.03,
			transactionType: Deposit,
		}
		var aTransfer Transaction = Transaction{
			date:            time.Now(),
			amount:          0.01,
			transactionType: Transfer,
		}
		transactions = append(transactions, aTransfer)
		transactions = append(transactions, aDeposit)
		// act
		var actual1 float32 = CalculateBalance1(transactions)
		var actual2 float32 = CalculateBalance2(transactions)
		// assert
		var expected float32 = 0.02
		AssertEqual(t, fmt.Sprintf("%f", actual1), fmt.Sprintf("%f", expected))
		AssertEqual(t, fmt.Sprintf("%f", actual2), fmt.Sprintf("%f", expected))
	})
}
