# 1 - Variables

## 1.1 Declaring variables


```go
// Expressive way
var walletHolder string = "Alberto Basalo"
// Implicit way
var walletHolder = "Alberto Basalo"
// Shorthand
walletHolder := "Alberto Basalo"
```


## 1.2 Primitive Types

```go
var walletHolder string = "Alberto Basalo"
// Attention to different kinds of numbers
var walletBalance float32 = 3.1416
var numberOfTransactions uint = 42
var isActive bool = true
var creationDate time.Time = time.Now()
```

## 1.3 Custom Types


### 1.3.1 Enums

```go
type TransactionType string
const (
	Deposit    TransactionType = "Deposit"
	Withdrawal TransactionType = "Withdrawal"
	Transfer   TransactionType = "Transfer"
)
```

### 1.3.2 Structs


```go
type Transaction struct {
	date            time.Time
	amount          float32
	transactionType TransactionType
}
var newTransaction Transaction = Transaction{
	date:            time.Now(),
	amount:          0.01,
	transactionType: Deposit, // last comma is mandatory
}
```

### 1.4 Arrays

```go
var transactions []Transaction
transactions = append(transactions, newTransaction)
```

### 1.5 Tests

```go
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
```


### [Back to index](https://github.com/AtomicBuilders/gluon/blob/main/docs/index.md)