package main

import (
	"fmt"
	"time"
)

func InitializeWallet() {
	fmt.Println("💰 My wallet")
	var walletHolder string = "Alberto Basalo"
	fmt.Println("Holder: " + walletHolder)
	var walletBalance int = 0
	fmt.Println("Balance:", walletBalance)
	numberOfTransactions := 0
	fmt.Println("Number of Transactions:", numberOfTransactions)
	const minimumTransactionAmount float32 = 0.01
	fmt.Println("Minimum Transaction Amount:", minimumTransactionAmount)
	var isActive bool = true
	fmt.Println("Is Active:", isActive)
	var creationDate time.Time = time.Now()
	const longTimeLayout = "2006-01-02 15:04:05"
	fmt.Println("Creation Date:", creationDate.Format(longTimeLayout))
	const oneYear = time.Hour * 24 * 365
	var expirationDate time.Time = creationDate.Add(oneYear * 2)
	const shortTimeLayout = "2006-01-02"
	fmt.Println("Expiration Date:", expirationDate.Format(shortTimeLayout))
}
