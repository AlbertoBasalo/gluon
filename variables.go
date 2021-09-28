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
	type CurrencyType string
	const (
		Bitcoin CurrencyType = "Bitcoin"
		Ether   CurrencyType = "Ether"
		USD     CurrencyType = "USD"
		Euro    CurrencyType = "Euro"
	)
	var walletCurrency CurrencyType = Ether
	fmt.Println("Currency:", walletCurrency)
	var numberOfTransactions uint = 0
	fmt.Println("Number of Transactions:", numberOfTransactions)
	const minimumTransactionAmount float32 = 0.01
	fmt.Println("Minimum Transaction Amount:", minimumTransactionAmount)
	var isActive bool = true
	fmt.Println("Is Active:", isActive)
	var creationDate time.Time = time.Now()
	const longTimeLayout string = "2006-01-02 15:04:05"
	fmt.Println("Creation Date:", creationDate.Format(longTimeLayout))
	const oneYear time.Duration = time.Hour * 24 * 365
	var expirationDate time.Time = creationDate.Add(oneYear * 2)
	const shortTimeLayout string = "2006-01-02"
	fmt.Println("Expiration Date:", expirationDate.Format(shortTimeLayout))
	type TransactionType string
	const (
		Deposit    TransactionType = "Deposit"
		Withdrawal TransactionType = "Withdrawal"
		Transfer   TransactionType = "Transfer"
	)
	type transaction struct {
		date            time.Time
		amount          float32
		transactionType TransactionType
	}
	var newTransaction transaction = transaction{
		date:            time.Now(),
		amount:          0.01,
		transactionType: Deposit,
	}
	fmt.Println("New Transaction:", newTransaction)
	var transactions []transaction
	transactions = append(transactions, newTransaction)
	fmt.Println("Transactions:", transactions)
}
