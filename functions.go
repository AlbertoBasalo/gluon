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
		// sends a copy of the value of variable balance
		// declares and initializes two variables with the function results
		if newBalance, err := getNewBalance(balance, transaction); err != nil {
			// the statement also checks if there is an error
			fmt.Println("💥 Not processed:", err)
		} else {
			// but balance is not mutated inside the function
			println("current balance:", balance)
			println("new balance:", newBalance)
			balance = newBalance
		}
	}
	return balance
}

// current is a value not a pointer (no effect on the original value)
// function with multiple return values
func getNewBalance(current float32, transaction Transaction) (float32, error) {
	if transaction.amount < 0 {
		// return values on guard failure
		return current, fmt.Errorf("Transaction amount %v cannot be negative", transaction.amount)
	}
	// changes to local variable named current
	// initialized with a copy of the caller value
	switch transaction.transactionType {
	case Deposit:
		current += transaction.amount
	case Withdrawal:
		fallthrough
	case Transfer:
		current -= transaction.amount
	}
	// return values on happy path
	return current, nil
}
