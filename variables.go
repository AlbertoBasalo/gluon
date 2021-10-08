package main

import (
	"fmt"
	"time"
) // several imports in one line

// A kind of string enum
type CurrencyType string

const (
	Bitcoin CurrencyType = "Bitcoin"
	Ether   CurrencyType = "Ether"
	USD     CurrencyType = "USD"
	Euro    CurrencyType = "Euro"
)

// structs for complex types
type Wallet struct {
	holder               string
	balance              float32
	iaActive             bool
	minimumTransaction   float32
	numberOfTransactions uint
}

type TransactionType string

const (
	Deposit    TransactionType = "Deposit"
	Withdrawal TransactionType = "Withdrawal"
	Transfer   TransactionType = "Transfer"
)

type Transaction struct {
	date            time.Time
	amount          float32
	transactionType TransactionType
}

func InitializeWallet() Wallet {
	fmt.Println("💰 My wallet")
	// standard way to declare, set type and assign value to a variable
	var walletHolder string = "Alberto Basalo"
	fmt.Println("Holder: " + walletHolder)
	// attention for different kinds of numbers
	var walletBalance float32 = 0
	fmt.Println("Balance:", walletBalance)
	var walletCurrency CurrencyType = Ether
	fmt.Println("Currency:", walletCurrency)
	var numberOfTransactions uint = 0
	fmt.Println("Number of Transactions:", numberOfTransactions)
	const minimumTransaction float32 = 0.01
	fmt.Println("Minimum Transaction:", minimumTransaction)
	var isActive bool = true
	fmt.Println("Is Active:", isActive)
	// date and time came on their own package
	var creationDate time.Time = time.Now()
	// formating a date string is intimidating the first time
	const longTimeLayout string = "2006-01-02 15:04:05"
	fmt.Println("Creation Date:", creationDate.Format(longTimeLayout))
	const oneYear time.Duration = time.Hour * 24 * 365
	var expirationDate time.Time = creationDate.Add(oneYear * 2)
	const shortTimeLayout string = "2006-01-02"
	fmt.Println("Expiration Date:", expirationDate.Format(shortTimeLayout))
	var newTransaction Transaction = Transaction{
		date:            time.Now(),
		amount:          0.01,
		transactionType: Deposit, // last comma is mandatory
	}
	fmt.Println("New Transaction:", newTransaction)
	// Arrays
	var transactions []Transaction
	transactions = append(transactions, newTransaction)
	fmt.Println("Transactions:", transactions)
	var myWallet Wallet = Wallet{
		holder:               walletHolder,
		balance:              walletBalance,
		iaActive:             isActive,
		minimumTransaction:   minimumTransaction,
		numberOfTransactions: numberOfTransactions,
	}
	return myWallet
}
