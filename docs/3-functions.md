# 3 - Functions

## 3.1 Simple


```go
// simple function
func ProcessTransactions(transactions []Transaction) {
	fmt.Println("Do stuff...")
}
ProcessTransactions(myTransactions)

```


## 3.2 Anonymous

```go
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
```

## 3.3 Multiple return values


### 3.3.1 Returning multiple values

```go
// function with multiple return values
func getNewBalance(current float32, transaction Transaction) (float32, error) {
	if transaction.amount < 0 {
		// return values on guard failure
		return current, fmt.Errorf("Transaction amount %v cannot be negative", transaction.amount)
	}
	var something float32 = 42
	// return values on happy path
	return something, nil
}
```

### 3.3.2 Calling and saving the results


```go
// declares and initializes two variables with the function results
if newBalance, err := getNewBalance(balance, transaction); err != nil {
	// the statement also checks if there is an error
	fmt.Println("💥 Not processed:", err)
} else {
	balance = newBalance
}
```


## 3.4 Values, not references

```go
var balance float32 = 0
// current is a value not a pointer (no effect on the original value)
func getNewBalance(current float32, transaction Transaction) float32 {
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
	return current, nil
}
// sends a copy of the value of variable balance
newBalance, err := getNewBalance(balance, transaction);
// but balance is not mutated inside the function
println("current balance:", balance)
println("new balance:", newBalance)
balance = newBalance

```

### [Back to index](https://github.com/AtomicBuilders/gluon/blob/main/docs/index.md)