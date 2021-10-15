package main

import "fmt"

// simple function
func ProcessTransactions(transactions []Transaction) {
	balance := CalculateBalance(transactions)
	// anonymous function saved in a variable
	var printWealth = func(wealth float32) {
		switch {
		case wealth < 0:
			fmt.Println("😖 You have no money")
		case wealth < 100:
			fmt.Println("😃 You have money")
		default:
			fmt.Println("🤑 You have a lot of money")
		}
	}
	// calls the anonymous function, by its variable name
	printWealth(balance)
}

func CalculateBalance(transactions []Transaction) float32 {
	var balance float32 = 0
	for _, transaction := range transactions {
		// declares and initializes two variables,
		if newBalance, err := getNewBalance(balance, transaction); err != nil {
			// also checks if there is an error
			fmt.Println("💥 Not processed:", err)
		} else {
			balance = newBalance
		}
	}
	return balance
}

// function with multiple return values
func getNewBalance(current float32, transaction Transaction) (float32, error) {
	if transaction.amount < 0 {
		// return values on guard failure
		return current, fmt.Errorf("Transaction amount %v cannot be negative", transaction.amount)
	}
	var newBalance float32 = 0.0
	switch transaction.transactionType {
	case Deposit:
		newBalance = current + transaction.amount
	case Withdrawal:
		fallthrough
	case Transfer:
		newBalance = current - transaction.amount
	}
	// return values on happy path
	return newBalance, nil
}
