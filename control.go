package main

import "fmt"

func CalculateBalance1(transactions []Transaction) float32 {
	var balance float32 = 0
	// Classic for loop with index
	for i := 0; i < len(transactions); i++ {
		var transaction = transactions[i]
		// simple conditionals
		if transaction.transactionType == Deposit {
			balance += transaction.amount
		} else {
			balance -= transaction.amount
		}
	}
	// nested if statements
	if balance < 0 {
		fmt.Println("😖 You have no money")
	} else if balance < 100 {
		fmt.Println("😃 You have money")
	} else {
		fmt.Println("🤑 You have a lot of money")
	}
	return balance
}

func CalculateBalance2(transactions []Transaction) float32 {
	var balance float32 = 0
	// For range loop can interate over slices
	for _, transaction := range transactions {
		// switch statement
		switch transaction.transactionType {
		case Deposit:
			balance += transaction.amount
		case Withdrawal:
			fallthrough
		case Transfer:
			balance -= transaction.amount
		}
	}
	// switch statement without initialisation
	switch {
	case balance < 0:
		fmt.Println("😖 You have no money")
	case balance < 100:
		fmt.Println("😃 You have money")
	default:
		fmt.Println("🤑 You have a lot of money")
	}
	return balance
}
