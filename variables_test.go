package main

import (
	"fmt"
	"testing"
)

func TestInitializeWallet(t *testing.T) {
	t.Run("it should initialize the wallet", func(t *testing.T) {
		// arrange
		// act
		actual := InitializeWallet()
		// assert
		expected := wallet{
			holder:               "Alberto Basalo",
			balance:              0,
			iaActive:             true,
			minimumTransaction:   0.01,
			numberOfTransactions: 0,
		}
		AssertEqual(t, fmt.Sprintf("%#v", &actual), fmt.Sprintf("%#v", &expected))
	})
}
