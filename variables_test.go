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
		expected := Wallet{
			holder:               "Alberto Basalo",
			balance:              0,
			iaActive:             true,
			minimumTransaction:   0.01,
			numberOfTransactions: 0,
		}
		// comparing structs is not that easy,
		// so we use the fmt package to compare
		// the string representation of the structs
		AssertEqual(t, fmt.Sprintf("%#v", &actual), fmt.Sprintf("%#v", &expected))
	})
}
